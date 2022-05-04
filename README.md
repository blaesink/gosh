# Gosh
![goVersion](https://badgen.net/badge/golang/1.18/cyan)
![status](https://badgen.net/badge/status/alpha?color=red)
![version](https://badgen.net/badge/version/0.0.1?color=green) <!-- Todo scrape this from some VERSION file later. -->

A **Go**od **Sh**ell for UNIX.

- "Oh my gosh, this is great!"
- "Oh my gosh, it's smooth to use!"
- "Oh my gosh, where has this been my entire life?"
## Notes
- Gosh came about from my desire to learn Go (and its core features) in a hands-on approach.
- That means that there will be things that will arouse hypertension in more experienced Gophers. Thanks for sticking with me as I figure them out!
- This means that gosh is *pre-pre-pre-alpha* and isn't intended for anyone to really use for a while. Unless you are willing to drive a car with two tires and no steering wheel for a while.
## Planned Features
Gosh aims to utilize modern technologies and feature sets from other shells.
Many of these are by means of unashamedly re-implementing certain features that
are "extensions" of said shells.

- **Asynchronicity** &mdash; The inspiration comes from [fish-async-prompt](https://github.com/acomagu/fish-async-prompt). While shells aren't heavy to run, having any noticeable "lag" degrades user experience. Allowing the user to tell the shell that a command should be run in a separate thread, and return back to them later as they do something else is not only "cool", but keeps the user's flow.
- **Intellisense** &mdash; In name only. Gosh will keep track of **how often** and **how recently** a user issued a command. Gosh will attempt to suggest that command as you type it (even if aliased!) to save keystrokes and time. Influence here comes from [rupa/z](https://github.com/rupa/z), which will also be included.
- **Batteries Included** &mdash; The base Gosh install strives to provide a set of **useful** and **positive** tools. Many of these take inspiration or are ports (denoted as such).
    - Z from [rupa/z](https://github.com/rupa/z). Keeps track of user directory switching and allows users to jump around via fuzzy searching. Incredible time saver and enhances user flow.
    - More to come!
- **Extensibility** &mdash; What's a shell without allowing you to modify it like Skyrim? Gosh provides an easy way to add user functionality.
    - **Easy Aliasing** &mdash; User aliases are separated from user commands and are defined in a YAML file for the sake of you, the user's, sanity.
    - **Streamlined Syntax and Scripting** &mdash;  Inspiration was taken from `fish` in that POSIX compliance is largely cumbersome for a single user. Gosh will provide a scripting language and syntax doesn't intend to conform to POSIX standards (but it will let you use those scripts and syntax!), human readability is the most important when trying to make something that *works*.
    - **Library Extensions** &mdash; Gosh (in the future) will allow users to add their own Go scripts and extensions by exporting its library to a separate repository. This will allow for deeper, more in-depth customizabilty outside of writing configs in Gosh's own scripting language.
