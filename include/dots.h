#ifndef _DOTS_H
#define _DOTS_H

typedef struct {
    char *name;
    char *source;
    char *target;
    char *hook;
} Dotfile;

typedef struct {
    Dotfile *dots;
    unsigned long count;
} Dotfiles;

Dotfiles dots_init();
int dots_link(Dotfile *dotfile);
int dots_unlink(Dotfile *dotfile);
void dots_destroy(Dotfiles *d);
void dots_init_dir(const char *dirname);

#endif
