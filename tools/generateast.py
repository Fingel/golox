#!/usr/local/bin/python3

import sys

TYPES = [
    'Binary     : left Expr,operator Token,right Expr',
    'Grouping   : expression Expr',
    'Literal    : value interface{}',
    'Unary      : operator Token,right Expr'
]

def define_ast(output_dir, base_name, types):
    path = f'{output_dir}/{base_name}.go'
    with open(path, 'w') as f:
        f.write('package lox\n\n')
        f.write(f'type {base_name} struct {{\n')
        f.write('}\n\n')

        for t in types:
            class_name = t.split(':')[0].strip()
            fields = t.split(':')[1].strip()
            define_type(f, base_name, class_name, fields)

def define_type(f, base_name, class_name, fields):
    f.write(f'type {class_name} struct {{\n')
    f.write(f'  {base_name}\n')
    for field in fields.split(','):
        f.write(f'  {field}\n')
    f.write('}\n\n')

def main():
    if len(sys.argv) < 2:
        print('Usage: generateast <output directory>')
        sys.exit(64)
    output_dir = sys.argv[1]
    define_ast(output_dir, 'Expr', TYPES)


if __name__ == '__main__':
    main()
