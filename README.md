# Recommended Software
These are my personal preferences which I suggest for those with a similar taste in minimalism and extensibility.

Operating Systems:
* Void Linux (glibc for 32-bit musl for 64-bit)
* OpenBSD
* Alpine Linux
* Hyperbola GNU/Linux

Shells:
* dash/Busybox's ash (should be the default shell regardless of login/user shell)
* es/Plan 9's rc (should NEVER be the default shell on Linux or BSD)
* OpenBSD's ksh and its forks/ports: loksh, mksh, pdksh
* fish (features are worth the bloat while using a terminal but not for scripts)

Coreutils (alternatives to GNU):
* Busybox
* Plan 9

libc (alternatives to glibc):
* musl-libc
* dietlibc
* uClibc

https://etalab.net/compare_libcs.html

C compilers (alternatives to gcc):
* tcc
* 9c

Init Systems (alternatives to SystemD/OpenRC):
* runit
* s6

Programming Languages:
* task automation: POSIX-compliant shell scripts *(rather than Python)*
* scientific computing: Go, C (with gnuplot), Julia, R *(rather than MatLab, Octave, numpy/scipy/matplotlib or whatever)*
* server-side: Go, nodejs, PHP *(rather than Ruby on Rails, Python, ASP)*
* general purpose: C, Go, Lua *(rather than Java, C++, Python)*
* regular expressions: Perl, awk, sed *(rather than Python)*

Terminal Emulators:
* Suckless's st
* xterm
* kitty (bloated but may be worth it if you have a lot of spare RAM)

Text Editors:
* neovim
* vis
* nvi
* sam
* joe

File Managers *(rather than ranger, lf, or just using the shell)*:
* vifm 
* nnn

Readers:
* zathura, apvlv (for DjVu, pdf, ps)
* vimpager or busybox's less (for text files or stdin)

Browsers:
* Firefox *(with heavy customization, i.e. Librefox)*
* GNU Icecat
* TorBrowser *(rather than using VPN services)*
* "Ungoogled" Chromium

Obscure browsers like surf and vimb are interesting but sub-par in terms of security. These are fine for viewing pages stored locally.

Browser Extenstions:
* uBlock Origin *(rather than NoScript, AdBlockPlus, etc.)*
* NoHTTP *(rather than HTTPS Everywhere)*
* CookieMasrer *(rather than Privacy Badger)*
* Tridactyl or Vim Vixen
* Dark Reader

Search Engines:
* https://searx.info
* https://duckduckgo.com/lite
