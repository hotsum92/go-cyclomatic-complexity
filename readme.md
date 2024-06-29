# go-cyclomatic-complexity

## cyclomatic complexity

M = E − N + 2P
M = 循環的複雑度
E = グラフのエッジ数
N = グラフのノード数
P = 連結されたコンポーネントの数

## gocyclo

[gocyclo](https://github.com/fzipp/gocyclo)
[gocycloを使ってgo言語のプロダクトをシンプルに維持する｜moli9ma](https://note.com/moli9ma/n/n7e4ade98c443)

### install

```bash
go install github.com/fzipp/gocyclo/cmd/gocyclo@latest
```

### usage

```bash
gocyclo .
gocyclo main.go
```

### output

<complexity> <package> <function> <file:line:column>

```
3 main switchcase main.go:17:1
```

## bugspots

[bugspots-go](https://github.com/masmgr/bugspots-go)

### run

```bash
go build .
go run bugspots-go
```
### usage

```bash
~/git/bugspots-go/bugspots-go .
```

## ast-walker

[ast-walker](https://github.com/nirasan/ast-walker)

[GoのAST全部見る](https://monpoke1.hatenablog.com/entry/2018/12/16/110943)

### run

```bash
go build .
go run ast-walker
```

### usage

```bash
~/git/ast-walker/ast-walker -type *ast.AssignStmt ./main.go | grep -E '^COMMAND' | wc -l
~/git/ast-walker/ast-walker ./main.go
```

## git

```
git diff --stat :/first HEAD
git log --stat :/first HEAD
git rebase --exec 'cat $(git ls-files) | wc -l' :/first
git rebase --exec "$(cat ./stat.sh | tr '\n' ';')" :/first
```


