#include "colors.h"
#include "parser.h"

#include <getopt.h>
#include <stdlib.h>
#include <string.h>

void print_help(void);
int dotfile_op(int (*fn)(Dotfile *dotfile));

#define VERSION "0.0.5 (dev)"

int main(int argc, char *argv[])
{
    while (1) {
        /* clang-format off */
        static struct option options[] = {
            {"init",    required_argument, 0, 'i'},
            {"link",    required_argument, 0, 'l'},
            {"unlink",  required_argument, 0, 'u'},
            {"help",    no_argument,       0, 'h'},
            {"version", no_argument,       0, 'v'},
            {0,         0,                 0,   0},
        };
        /* clang-format on */
        int option_index = 0;
        int return_code  = 0;

        int c = getopt_long(argc, argv, "i:l:u:hv", options, &option_index);

        if (c == -1)
            break;

        switch (c) {
            case 0:
                break;

            case 'i':
                dots_init_dir(optarg);
                printf_succ("Created %sdotfile directory `%s`", RESET, optarg);
                break;
            case 'l':
                return_code = dotfile_op(dots_link);
                if (return_code == 0)
                    printf_succ("Finished%s. %s linked", RESET, optarg);
                else if (return_code == 2)
                    printf_err("ERROR%s: %s not found", RESET, optarg);

                return return_code;
                break;
            case 'u':
                return_code = dotfile_op(dots_unlink);
                if (return_code == 0)
                    printf_succ("Finished%s. %s unlinked", RESET, optarg);
                else if (return_code == 2)
                    printf_err("ERROR%s: %s not found", RESET, optarg);

                return return_code;
                break;
            case 'v':
                printf_succ("sween%s %s", RESET, VERSION);
                break;
            case '?':
            case 'h':
                print_help();
                exit(EXIT_SUCCESS);
                break;
            default:
                abort();
        }
    }

    return 0;
}

int dotfile_op(int (*fn)(Dotfile *dotfile))
{
    Dotfiles dotfiles = dots_init();
    int code          = 2; /* 2 means not found */

    size_t optarg_len = strlen(optarg) - 1;
    if (optarg[optarg_len] == '/')
        optarg[optarg_len] = '\0';

    if (strcmp(optarg, "all") == 0)
        for (size_t i = 0; i < dotfiles.count; i++)
            code = fn(&dotfiles.dots[i]);
    else
        for (size_t i = 0; i < dotfiles.count; i++) {
            if (strcmp(dotfiles.dots[i].name, optarg) == 0)
                code = fn(&dotfiles.dots[i]);
            else
                continue;
        }
    dots_destroy(&dotfiles);

    return code;
}

void print_help(void)
{
    printf("sween (Dotfile manager) version %s\n\n", VERSION);
    printf("USAGE:\n\n");
    printf("\tsween [OPTION] dotfile/all\n\n");
    printf("OPTIONS:\n\n");
    printf("\t-i DIR, --init   DIR  Create a new dotfiles directory\n");
    printf("\t-l DOT, --link   DOT  Link a dotfile to destination point\n");
    printf("\t-u DOT, --unlink DOT  Unlink a dotfile\n\n");
    printf("\t-v, --version         Prints version and exit\n");
    printf("\t-h, --help            Show this help\n");
    printf("\nHome page: MY GITHUB HERE\n");
}
