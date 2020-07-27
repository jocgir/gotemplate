---
title: CLI Usage
weight: 1
---
<!-- markdownlint-disable MD025 --->
# CLI Usage

```text
usage: ./gotemplate [<flags>] <command> [<args> ...]

An extended template processor for go.

See: https://coveo.github.io/gotemplate for complete documentation.

Flags:
  -h, --help                     Show context-sensitive help (also try --help-man). or set _GOTEMPLATE_HELP
      --color                    Force rendering of colors event if output is redirected or set
                                 _GOTEMPLATE_COLOR (alias: --c)
  -v, --version                  Get the current version of gotemplate
      --template-log-level=level  
                                 Set the template logging level. Accepted values: disabled, panic, fatal,
                                 error, warning, info, debug, trace or set _GOTEMPLATE_TEMPLATE_LOG_LEVEL
                                 (alias: --tll)
  -L, --internal-log-level=level  
                                 Set the internal logging level. Accepted values: disabled, panic, fatal,
                                 error, warning, info, debug, trace or set _GOTEMPLATE_INTERNAL_LOG_LEVEL
                                 (alias: --ill, --ll, --log-level)
  -F, --internal-log-file-path=path  
                                 Set a file where verbose logs should be written or set
                                 _GOTEMPLATE_INTERNAL_LOG_FILE_PATH (alias: --ilfp)
      --template-log-file-level=level  
                                 Set the template logging level for the verbose logs file or set
                                 _GOTEMPLATE_TEMPLATE_LOG_FILE_LEVEL (alias: --tlfl)
      --internal-log-file-level=level  
                                 Set the internal logging level for the verbose logs file or set
                                 _GOTEMPLATE_INTERNAL_LOG_FILE_LEVEL (alias: --ilfl)
      --base                     Turn off all addons (they could then be enabled explicitly) or set
                                 _GOTEMPLATE_BASE
      --razor                    Razor Addon (ON by default) or set _GOTEMPLATE_RAZOR (off: --no-razor)
      --extension                Extension Addon (ON by default) or set _GOTEMPLATE_EXTENSION (alias: --ext)
                                 (off: --next, --no-ext, --no-extension)
      --math                     Math Addon (ON by default) or set _GOTEMPLATE_MATH (off: --no-math)
      --sprig                    Sprig Addon (ON by default) or set _GOTEMPLATE_SPRIG (off: --no-sprig)
      --data                     Data Addon (ON by default) or set _GOTEMPLATE_DATA (off: --no-data)
      --logging                  Logging Addon (ON by default) or set _GOTEMPLATE_LOGGING (off:
                                 --no-logging)
      --runtime                  Runtime Addon (ON by default) or set _GOTEMPLATE_RUNTIME (off:
                                 --no-runtime)
      --utils                    Utils Addon (ON by default) or set _GOTEMPLATE_UTILS (off: --no-utils)
      --net                      Net Addon (ON by default) or set _GOTEMPLATE_NET (off: --nnet, --no-net)
      --os                       OS Addon (ON by default) or set _GOTEMPLATE_OS (off: --no-os, --nos)
      --git                      Git Addon (ON by default) or set _GOTEMPLATE_GIT (off: --ngit, --no-git)

Args:
  [<templates>]  Template files or commands to process

Commands:
  help [<command>...]
    Show help.


  run [<flags>] [<templates>...]

        --delimiters={{,}},@       Define the default delimiters for go template (separate the left, right
                                   and razor delimiters by a comma) or set _GOTEMPLATE_DELIMITERS (alias:
                                   --d, --del)
    -i, --import=file ...          Import variables files (could be any of YAML, JSON or HCL format) or set
                                   _GOTEMPLATE_IMPORT
        --import-if-exist=file ...  
                                   Import variables files (do not consider missing file as an error) or set
                                   _GOTEMPLATE_IMPORT_IF_EXIST (alias: --iie)
    -V, --var=values ...           Import named variables (if value is a file, the content is loaded) or set
                                   _GOTEMPLATE_VAR
    -t, --type=TYPE                Force the type used for the main context (Json, Yaml, Hcl) or set
                                   _GOTEMPLATE_TYPE
    -p, --patterns=pattern ...     Additional patterns that should be processed by gotemplate or set
                                   _GOTEMPLATE_PATTERNS
    -e, --exclude=pattern ...      Exclude file patterns (comma separated) when applying gotemplate
                                   recursively or set _GOTEMPLATE_EXCLUDE
    -o, --overwrite                Overwrite file instead of renaming them if they exist (required only if
                                   source folder is the same as the target folder) or set
                                   _GOTEMPLATE_OVERWRITE
    -s, --substitute=exp ...       Substitute text in the processed files by applying the regex substitute
                                   expression (format: /regex/substitution, the first character acts as
                                   separator like in sed, see: Go regexp) or set _GOTEMPLATE_SUBSTITUTE
    -E, --remove-empty-lines       Remove empty lines from the result or set _GOTEMPLATE_REMOVE_EMPTY_LINES
                                   (alias: --re, --rel, --remove-empty)
    -r, --recursive                Process all template files recursively or set _GOTEMPLATE_RECURSIVE
    -R, --recursion-depth=depth    Process template files recursively specifying depth or set
                                   _GOTEMPLATE_RECURSION_DEPTH (alias: --rd)
        --source=folder            Specify a source folder (default to the current folder) or set
                                   _GOTEMPLATE_SOURCE (alias: --s)
        --target=folder            Specify a target folder (default to source folder) or set
                                   _GOTEMPLATE_TARGET (alias: --t)
    -I, --stdin                    Force read of the standard input to get a template definition (useful
                                   only if GOTEMPLATE_NO_STDIN is set) or set _GOTEMPLATE_STDIN
    -f, --follow-symlinks          Follow the symbolic links while using the recursive option or set
                                   _GOTEMPLATE_FOLLOW_SYMLINKS (alias: --fs)
    -P, --print                    Output the result directly to stdout or set _GOTEMPLATE_PRINT
    -d, --disable                  Disable go template rendering (used to view razor conversion) or set
                                   _GOTEMPLATE_DISABLE
        --accept-no-value          Do not consider rendering <no value> as an error or set
                                   GOTEMPLATE_NO_VALUE (alias: --anv, --no-value, --nv)
    -S, --strict-error-validation  Consider error encountered in any file as real error or set
                                   GOTEMPLATE_STRICT_ERROR (alias: --sev, --strict)
        --strict-assignations-validation=warning  
                                   Enforce strict assignation validation on global variables or set
                                   _GOTEMPLATE_STRICT_ASSIGNATIONS_VALIDATION (alias: --sav)
        --ignore-missing-import    Exit with code 0 even if import does not exist or set
                                   _GOTEMPLATE_IGNORE_MISSING_IMPORT (alias: --imi)
        --ignore-missing-source    Exit with code 0 even if source does not exist or set
                                   _GOTEMPLATE_IGNORE_MISSING_SOURCE (alias: --ims)
        --ignore-missing-paths     Exit with code 0 even if import or source do not exist or set
                                   _GOTEMPLATE_IGNORE_MISSING_PATHS (alias: --imp)

  list [<flags>] [<filters>...]
    Get detailed help on gotemplate functions

    -f, --functions  Get detailed help on function
    -t, --templates  List the available templates
    -l, --long       Get detailed list
    -a, --all        List all
    -c, --category   Group functions by category
```
