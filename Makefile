SHELL:=bash

.ONESHELL: # It all runs on one single shell

MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-biultin-rules

dev_env:=$(if $(DEV_ENV), $(DEV_ENV), 'docker')
app_slug:=finance

.DEFAULT_GOAL := run # $ make runs the application by default
