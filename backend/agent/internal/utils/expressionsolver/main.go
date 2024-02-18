package expressionsolver

import (
	"log"
	"math"
	"strconv"
	"strings"
	"time"
)

func EvaluateExpression(expression string, ops map[string]int) float64 {
    // Приведение выражения к обратной польской нотации (ОПН)
    rpnExpression := toRPN(expression)

    // Вычисление значения выражения из ОПН
    res := evaluateRPN(rpnExpression, ops)
    return roundFloat(res, 4)
}

// Преобразование выражения в обратную польскую нотацию (ОПН)
func toRPN(expression string) []string {
    var rpn []string
    var operators []string

    tokens := tokenize(expression)

    for _, token := range tokens {
        if isOperator(token) {
            for len(operators) > 0 && precedence(operators[len(operators)-1]) >= precedence(token) {
                rpn = append(rpn, operators[len(operators)-1])
                operators = operators[:len(operators)-1]
            }
            operators = append(operators, token)
        } else if token == "(" {
            operators = append(operators, token)
        } else if token == ")" {
            for operators[len(operators)-1] != "(" {
                rpn = append(rpn, operators[len(operators)-1])
                operators = operators[:len(operators)-1]
            }
            operators = operators[:len(operators)-1]
        } else {
            rpn = append(rpn, token)
        }
    }

    for len(operators) > 0 {
        rpn = append(rpn, operators[len(operators)-1])
        operators = operators[:len(operators)-1]
    }

    return rpn
}

// Вычисление значения выражения из обратной польской нотации (ОПН)
func evaluateRPN(rpn []string, ops map[string]int) float64 {
    var stack []float64

    for _, token := range rpn {
        if isOperator(token) {
            operand2 := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            operand1 := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            result := operate(token, operand1, operand2, ops)
            stack = append(stack, result)
        } else {
            num, _ := strconv.ParseFloat(token, 64)
            stack = append(stack, num)
        }
    }

    return stack[0]
}

func isOperator(op string) bool {
    return op == "+" || op == "-" || op == "*" || op == "/"
}

func precedence(op string) int {
    switch op {
    case "+", "-":
        return 1
    case "*", "/":
        return 2
    }
    return 0
}

func operate(operator string, operand1, operand2 float64, ops map[string]int) float64 {
    switch operator {
    case "+":
		res := operand1 + operand2
		time.Sleep(time.Second*time.Duration(ops["+"]))
        return res
    case "-":
		res := operand1 - operand2
		time.Sleep(time.Second*time.Duration(ops["-"]))
        return res
    case "*":
		res := operand1 * operand2
		time.Sleep(time.Second*time.Duration(ops["*"]))
        return res
    case "/":
		res := operand1 / operand2
		time.Sleep(time.Second*time.Duration(ops["/"]))
        return res
	default:
		log.Panicf("Unexpected operator: %s", operator)
		return 0	
	}	
}

func tokenize(expression string) []string {
    expression = strings.ReplaceAll(expression, "–", "-")
	var tokens []string
	currentToken := ""
	negativeSign := "-"

	for i, char := range expression {
		token := string(char)
		if token == "+" || token == "-" || token == "*" || token == "/" || token == "(" || token == ")" {
			if currentToken != "" {
				tokens = append(tokens, currentToken)
				currentToken = ""
			}
			// Проверка, является ли знак "-" отрицательным числом или оператором
			if token == "-" && (i == 0 || expression[i-1:i] == "(") {
				currentToken += negativeSign
			} else {
				tokens = append(tokens, token)
			}
		} else if token == " " { // Изменение токенизации: учесть пробелы
			continue
		} else {
			currentToken += token
		}
	}

	if currentToken != "" {
		tokens = append(tokens, currentToken)
	}

	return tokens
}

func roundFloat(val float64, precision uint) float64 {
    ratio := math.Pow(10, float64(precision))
    return math.Round(val*ratio) / ratio
}