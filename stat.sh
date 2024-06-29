#!/bin/bash

files="./main.go"

FuncDecl=$(~/git/ast-walker/ast-walker -type *ast.FuncDecl $files | grep -E '^COMMAND' | wc -l)
StructType=$(~/git/ast-walker/ast-walker -type *ast.StructType $files | grep -E '^COMMAND' | wc -l)
InterfaceType=$(~/git/ast-walker/ast-walker -type *ast.InterfaceType $files | grep -E '^COMMAND' | wc -l)
IfStmt=$(~/git/ast-walker/ast-walker -type *ast.IfStmt $files | grep -E '^COMMAND' | wc -l)
SwitchStmt=$(~/git/ast-walker/ast-walker -type *ast.SwitchStmt $files | grep -E '^COMMAND' | wc -l)
GitLog=$(git log --no-decorate --oneline HEAD^..HEAD)

echo "$GitLog:FuncDecl=$FuncDecl StructType=$StructType InterfaceType=$InterfaceType IfStmt=$IfStmt SwitchStmt=$SwitchStmt"
