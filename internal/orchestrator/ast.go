package orchestrator

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

//ASTNode представляет узел абстрактного синтаксического дерева (AST)
type ASTNode struct {
	IsLeaf        bool     //флаг, указывающий, является ли узел листом (числом)
	Value         float64  //числовое значение для листовых узлов
	Operator      string   //оператор для внутренних узлов
	Left, Right   *ASTNode //левый и правый дочерние узлы
	TaskScheduled bool     //флаг, указывающий, запланирована ли задача для этого узла
}

//ParseAST преобразует строковое выражение в синтаксическое дерево
//возвращает корневой узел AST или ошибку при невалидном выражении
func ParseAST(expression string) (*ASTNode, error) {
	//удаляем все пробелы из выражения
	expr := strings.ReplaceAll(expression, " ", "")
	if expr == "" {
		return nil, fmt.Errorf("пустое выражение")
	}

	p := &parser{input: expr, pos: 0}
	node, err := p.parseExpression()
	if err != nil {
		return nil, err
	}

	//проверяем, что весь ввод был обработан
	if p.pos < len(p.input) {
		return nil, fmt.Errorf("неожиданный токен на месте %d", p.pos)
	}

	return node, nil
}

//parser реализует рекурсивный нисходящий парсер
type parser struct {
	input string //входная строка для парсинга
	pos   int    //текущая позиция в строке
}

//peek возвращает текущий символ без перемещения позиции
func (p *parser) peek() rune {
	if p.pos < len(p.input) {
		return rune(p.input[p.pos])
	}
	return 0
}

//get возвращает текущий символ и перемещает позицию вперед
func (p *parser) get() rune {
	ch := p.peek()
	if ch != 0 {
		p.pos++
	}
	return ch
}

//parseExpression обрабатывает сложение и вычитание
func (p *parser) parseExpression() (*ASTNode, error) {
	node, err := p.parseTerm()
	if err != nil {
		return nil, err
	}

	//обрабатываем цепочку операций + и -
	for {
		ch := p.peek()
		if ch == '+' || ch == '-' {
			op := string(p.get())
			right, err := p.parseTerm()
			if err != nil {
				return nil, err
			}
			node = &ASTNode{
				Operator: op,
				Left:     node,
				Right:    right,
			}
		} else {
			break
		}
	}

	return node, nil
}

//parseTerm обрабатывает умножение и деление
func (p *parser) parseTerm() (*ASTNode, error) {
	//парсим первый фактор
	node, err := p.parseFactor()
	if err != nil {
		return nil, err
	}

	//обрабатываем цепочку операций * и /
	for {
		ch := p.peek()
		if ch == '*' || ch == '/' {
			op := string(p.get())
			right, err := p.parseFactor()
			if err != nil {
				return nil, err
			}
			node = &ASTNode{
				Operator: op,
				Left:     node,
				Right:    right,
			}
		} else {
			break
		}
	}

	return node, nil
}

//parseFactor обрабатывает числа и выражения в скобках
func (p *parser) parseFactor() (*ASTNode, error) {
	ch := p.peek()

	//обработка скобок
	if ch == '(' {
		p.get() //пропускаем '('
		node, err := p.parseExpression()
		if err != nil {
			return nil, err
		}
		if p.peek() != ')' {
			return nil, fmt.Errorf("нет закрывающей скобки")
		}
		p.get() //пропускаем ')'
		return node, nil
	}
	//обработка чисел
	start := p.pos
	if ch == '+' || ch == '-' {
		p.get()
	}
	//собираем все цифры и точки
	for {
		ch = p.peek()
		if unicode.IsDigit(ch) || ch == '.' {
			p.get()
		} else {
			break
		}
	}

	//извлекаем токен числа
	token := p.input[start:p.pos]
	if token == "" {
		return nil, fmt.Errorf("неожиданное число на месте %d", start)
	}
	//парсим число в float64
	value, err := strconv.ParseFloat(token, 64)
	if err != nil {
		return nil, fmt.Errorf("невалидный номер %s", token)
	}
	return &ASTNode{
		IsLeaf: true,
		Value:  value,
	}, nil
}
