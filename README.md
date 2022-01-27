---
title: golang gin docker sample project
---

## 環境

実行には下記のものが必要

1. go
2. make (GNU Make 4.3で確認)
3. docker
4. [direnv](https://direnv.net/)
5. curl
6. github account

## 概要

1. golang + gin
2. docker
3. test
4. coverage
5. GitHub container registry
6. vscode devcontainer

## GitHub container registry

これでGitHubの自分のアカウントのPackagesを使うには下記の設定をする

1. PATを作成する。[Creating a personal access token](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token)
2. .envrc.local.template を .envrc.local にコピーし、CR_USER に、**GitHubユーザー名**を、CR_PATに作成した**PAT**を入れる。

※ .envrc.localの記述は、direnv が、.envrc を読んだ時に反映されるので、.envrc.local を変更した後は、`direnv reload` を実行する。

VSCodeのdevcontainerで開いて、Makefileがあるディレクトリでターミナルを開く。設定ができていれば、下記コマンドで、イメージがビルドされてpushされる。

```sh
make ghcr-login build push
```

## TODO

説明を追加する
