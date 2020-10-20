#ifndef _COLORS_H
#define _COLORS_H

#define RED   "\x1B[31m"
#define GRN   "\x1B[32m"
#define YEL   "\x1B[33m"
#define BLU   "\x1B[34m"
#define MAG   "\x1B[35m"
#define CYN   "\x1B[36m"
#define WHT   "\x1B[37m"
#define RESET "\x1B[0m"

/* Printf with '\t' at the beginning and '\n' at the end. And in green color. */
#define printf_succ(f_, ...)                                                                                \
    do {                                                                                                    \
        putchar('\t');                                                                                      \
        printf(GRN f_, __VA_ARGS__);                                                                        \
        putchar('\n');                                                                                      \
    } while (0)

/* Fprintf with '\t' at the beginning and '\n' at the end. And in red color. */
#define printf_err(f_, ...)                                                                                 \
    do {                                                                                                    \
        putc('\t', stderr);                                                                                 \
        fprintf(stderr, RED f_, __VA_ARGS__);                                                               \
        putc('\n', stderr);                                                                                 \
    } while (0)

#define check_oom(ptr)                                                                                      \
    do {                                                                                                    \
        if (!ptr) {                                                                                         \
            printf_err("ERROR%s: out of memory", RESET);                                                    \
            exit(EXIT_FAILURE);                                                                             \
        }                                                                                                   \
    } while (0)

#endif
