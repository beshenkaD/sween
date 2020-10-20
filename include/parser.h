#ifndef _PARSER_H
#define _PARSER_H

#include "dots.h"
#include "toml.h"

void parser_destroy_root(toml_table_t *manager);
toml_table_t *parser_load_root(void);
Dotfiles parser_get_dotfiles(toml_table_t *manager);
char *parser_get_dotfiles_dir(toml_table_t *manager);

#endif
