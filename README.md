# ai-naming
OpenAIのAPIを使用してプログラミングの命名をしてもらう

![demo](https://user-images.githubusercontent.com/70571576/222914576-6e764b9b-ae75-492c-acf6-44e57dedc5da.gif)

# get ApiKey
[Account API Keys - OpenAI API](https://platform.openai.com/account/api-keys)

# setup

```
$ cp .env.template .env
$ vi .env
$ make build-linux
```

# how to use

```
$ naming
```

```
? Choose a type:  [Use arrows to move, type to filter]
> 関数
  変数
```

```
? Enter an overview: 処理の内容を具体的に記述する
```

