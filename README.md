# gpt-naming

Have them name their programming using GPT-3.

![demo](https://user-images.githubusercontent.com/70571576/222914576-6e764b9b-ae75-492c-acf6-44e57dedc5da.gif)

## get ApiKey

[Account API Keys - OpenAI API](https://platform.openai.com/account/api-keys)

## setup

```shell
cp .env.template .env
vi .env
make build-linux
```

## how to use

```shell
naming
```

```shell
? Choose a type:  [Use arrows to move, type to filter]
  enum
  event
  exception
> function
  interface
  method
  namespace
```

```shell
? Enter an overview: Describe the process in detail
```

```shell
# Output
1. naming: reason.
```
