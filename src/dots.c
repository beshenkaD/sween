#include "dots.h"
#include "colors.h"
#include "parser.h"

#include <errno.h>
#include <stdlib.h>
#include <string.h>
#include <sys/stat.h>
#include <unistd.h>

#define MANAGER_NAME "manager.toml"

void dots_init_dir(const char *dirname)
{
    int code = mkdir(dirname, 0777);

    if (code == -1) {
        printf_err("ERROR%s: %s", RESET, strerror(errno));
        exit(EXIT_FAILURE);
    }

    char *path = calloc(1, strlen(dirname) + strlen(MANAGER_NAME) + 2);
    check_oom(path);

    sprintf(path, "%s/%s", dirname, MANAGER_NAME);

    FILE *manager = fopen(path, "w+");

    fprintf(manager, "[settings]\n");
    fprintf(manager, "dotfiles_dir = \"%s\"\n\n", dirname);
    fprintf(manager, "[cfg]\n");

    free(path);
    fclose(manager);
}

Dotfiles dots_init()
{
    toml_table_t *root = parser_load_root();

    if (!root) {
        printf_err("`%s` %snot found. Use --init option to create it.", MANAGER_NAME, RESET);
        exit(EXIT_FAILURE);
    }

    Dotfiles dots = parser_get_dotfiles(root);
    parser_destroy_root(root);

    return dots;
}

void dots_destroy(Dotfiles *d)
{
    for (size_t i = 0; i < d->count; i++) {
        if (d->dots[i].source)
            free(d->dots[i].source);
        if (d->dots[i].target)
            free(d->dots[i].target);
        if (d->dots[i].name)
            free(d->dots[i].name);
        if (d->dots[i].hook)
            free(d->dots[i].hook);
    }
    if (d->dots)
        free(d->dots);
}

static char *get_user()
{
    char *user = getenv("USER");
    if (strcmp(user, "root") == 0) {
        char *sudoer = getenv("SUDO_USER");
        if (sudoer)
            user = sudoer;
    }

    return user;
}

typedef enum {
    SOURCE,
    TARGET,
} PathType;

static char *get_canonical_path(char *path, PathType type)
{
    char *ret = NULL;

    toml_table_t *root = parser_load_root();
    char *dotfiles_dir = parser_get_dotfiles_dir(root);
    char *user         = get_user();

    size_t def_size = strlen(path) + strlen(user) + 1 + strlen("/home/") + 1;

    if (path[0] == '/') {
        ret = calloc(1, strlen(path) + 1);
        check_oom(ret);

        strcpy(ret, path);
    }
    else if (path[0] == '~') {
        if (type == TARGET)
            ret = calloc(1, def_size - 1);
        else
            ret = calloc(1, def_size + strlen(dotfiles_dir) + 1);
        check_oom(ret);

        char *new_path = malloc(strlen(path));
        check_oom(new_path);

        memmove(new_path, path + 1, strlen(path));

        if (type == TARGET)
            sprintf(ret, "/home/%s/%s", user, new_path);
        else
            sprintf(ret, "/home/%s/%s%s", user, dotfiles_dir, new_path);

        free(new_path);
    }
    else {
        if (type == TARGET) {
            ret = calloc(1, def_size);
            check_oom(ret);
            sprintf(ret, "/home/%s/%s", user, path);
        }
        else {
            ret = calloc(1, def_size + strlen(dotfiles_dir) + 1);
            check_oom(ret);
            sprintf(ret, "/home/%s/%s/%s", user, dotfiles_dir, path);
        }
    }

    parser_destroy_root(root);
    free(dotfiles_dir);

    return ret;
}

static char *strremove(char *str, const char *sub)
{
    char *p, *q, *r;
    if ((q = r = strstr(str, sub)) != NULL) {
        size_t len = strlen(sub);
        while ((r = strstr(p = r + len, sub)) != NULL) {
            memmove(q, p, r - p);
            q += r - p;
        }
        memmove(q, p, strlen(p) + 1);
    }
    return str;
}

int dots_link(Dotfile *dotfile)
{
    char *target = get_canonical_path(dotfile->target, TARGET);
    char *source = get_canonical_path(dotfile->source, SOURCE);

    int code = symlink(source, target);

    if (code == -1) {
        if (errno != EEXIST) {
            printf_err("ERROR%s: %s", RESET, strerror(errno));
        }
        else {
            code = 0;
        }
    }

    if (dotfile->hook) {
        char *hook = strremove(dotfile->hook, "\\");
        system(hook);
    }

    free(target);
    free(source);
    return code;
}

int dots_unlink(Dotfile *dotfile)
{
    char *target = get_canonical_path(dotfile->target, TARGET);
    int code     = unlink(target);

    if (code == -1) {
        if (errno == ENOENT)
            printf_err("ERROR%s: %s`%s`%s is not installed", RESET, GRN, dotfile->name, RESET);
        else
            printf_err("ERROR%s: %s", RESET, strerror(errno));
    }

    free(target);
    return code;
}
