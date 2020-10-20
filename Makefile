BIN     := sween
SRC_DIR := src
INC_DIR := include

CC		:= gcc
CFLAGS  := -std=gnu99 -Wall -Wextra -g -I$(INC_DIR)/ #-fanalyzer
LDFLAGS := 
LDLIBS  :=

PREFIX  := /usr/local/bin

SRC     := $(SRC_DIR)/main.c     \
		   $(SRC_DIR)/toml.c     \
		   $(SRC_DIR)/parser.c   \
		   $(SRC_DIR)/dots.c     \

HEADERS := $(INC_DIR)/toml.h     \
		   $(INC_DIR)/parser.h   \
		   $(INC_DIR)/colors.h   \
		   $(INC_DIR)/dots.h	 \

all: $(BIN)

release: CFLAGS   = -std=gnu99 -Ofast -flto -march=native -I$(INC_DIR)/
release: LDFLAGS  = -flto
release: all

$(BIN): $(SRC) $(HEADERS)
	$(CC) $(CFLAGS) $(SRC) $(LDFLAGS) -o $@ $(LDLIBS)

clean:
	$(RM) $(BIN)

install: $(BIN)
	mv $(BIN) $(PREFIX)

uninstall:
	rm $(PREFIX)/$(BIN)

.PHONY: all clean install uninstall
