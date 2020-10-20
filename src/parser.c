#include "parser.h"
#include "colors.h"
#include "dots.h"

#include <assert.h>
#include <stdlib.h>
#include <string.h>

toml_table_t *parser_load_root(void)
{
    FILE *manager = fopen("manager.toml", "r");
    if (!manager) {
        return NULL;
    }

    char errbuf[200];
    toml_table_t *root = toml_parse_file(manager, errbuf, sizeof(errbuf));

    fclose(manager);

    if (!root) {
        printf_err("ERROR%s: %s", RESET, errbuf);
        exit(1);
    }

    return root;
}

inline void parser_destroy_root(toml_table_t *manager) { toml_free(manager); }

char *parser_get_dotfiles_dir(toml_table_t *manager)
{
    toml_table_t *settings = toml_table_in(manager, "settings");
    if (!settings) {
        printf_err("ERROR%s: [settings] is missing. Use --init option to generate dotfiles directory",
                   RESET);
        exit(EXIT_FAILURE);
    }

    toml_raw_t dir_raw = toml_raw_in(settings, "dotfiles_dir");

    char *dotfiles_dir;
    toml_rtos(dir_raw, &dotfiles_dir);

    return dotfiles_dir;
}

Dotfiles parser_get_dotfiles(toml_table_t *manager)
{
    toml_table_t *dotfiles = toml_table_in(manager, "cfg");
    assert(dotfiles);

    size_t dotfiles_count = toml_table_ntab(dotfiles);

    Dotfiles ret;
    ret.dots = calloc(dotfiles_count, sizeof(Dotfile));
    check_oom(ret.dots);

    ret.count = dotfiles_count;

    for (size_t i = 0; i < dotfiles_count; i++) {
        const char *table_key = toml_key_in(dotfiles, i);
        toml_table_t *dot     = toml_table_in(dotfiles, table_key);

        toml_raw_t hook_raw   = toml_raw_in(dot, "hook");
        toml_raw_t target_raw = toml_raw_in(dot, "target");
        toml_raw_t source_raw = toml_raw_in(dot, "source");

        char *src, *target, *hook;
        toml_rtos(source_raw, &src);
        toml_rtos(target_raw, &target);
        toml_rtos(hook_raw, &hook);

        char *name = malloc(strlen(table_key) + 1);
        check_oom(name);

        strcpy(name, table_key);

        ret.dots[i].name   = name;
        ret.dots[i].source = src;
        ret.dots[i].target = target;
        ret.dots[i].hook   = hook;
    }

    return ret;
}
