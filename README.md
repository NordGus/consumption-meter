# ROM Stack

A Simple scaffold template to start building web apps with Go, HTMX and Web Components.

> _I have no idea what I'm doing_ - NordGus

## Why ROM Stack?

I discovered Go via a podcast talking about [Buffalo](https://gobuffalo.io), since then I was
fascinated with the possibility of shipping all the required data (files, templates, images,
assets, etc.) inside a single Go executable. To the point that one of my first experiments was a
[tool](https://github.com/NordGus/anguler) to embed an Angular application inside a Go executable.

It reminded me of the old cartridge-based game consoles and their ROMs from my childhood. So that's
why I decided to call it the ROM Stack, a Tech Stack designed to build ROM-like applications with Go.

## Did I forget to include batteries?

This is a simple setup, I've designed it to help me reduce friction between
idea and execution for my side-projects.

As I continue to work with it, I will update it so is not as bare bones or better
setup for serious development.

At the moment, this template includes a `devcontainer` setup for vscode that _works_(tm).
And also has a `nvm` setup to work on a local environment.

> But you're right. I completely forgot.

## So I used the template, now what?

### if you are working with the `devcontainer` setup

1. Install the tools needed by the Go extension.
2. Run `go mod edit <module name>` to rename the Go module.
3. Run `go mod tidy` to update Go dependencies.
4. Rename the project inside the `package.json`.
5. Run `yarn install` to update your `yarn.lock`.

### If you are working locally

1. Run `go mod edit <module name>` to rename the Go module.
2. Rename the project inside the `package.json`.
3. Run `go mod download` to download Go dependencies.
4. Run `yarn install` to download Node dependencies.

## Roadmap

- [x] Make the proof-of-concept.
- [x] Write my own manifesto-like `README.md`.
- [x] Learn how to set up `eslint`.
- [x] Setup `eslint`.
- [x] Learn how to set up `Prettier`.
- [x] Setup `Prettier`.
- [ ] Setup test environment.
- [ ] Learn how to use `Makefiles` for development environment.
- [ ] Setup `Makefile` to handle the development environment.
- [ ] Setup `Docker` image to build and serve the project.
- [ ] Learn how to set up a CI/CD pipeline for the project.
- [ ] Stop sucking at code.
- [ ] Stop procrastinating.
- [ ] Stop making checklist I can't complete.

## Important

**I have no idea what I'm doing**, so I do not recommend use this template for
production and any-scale deployments beyond personal education without heavy
modification.

If you have more experience setting up the missing pieces, or you consider
that I'm missing something else or misconfigured something. Don't be afraid
of collaborating. Open a PR. And please explain in details why and what it does.
Share your knowledge with other developers.

**Let's build cool stuff together**.
