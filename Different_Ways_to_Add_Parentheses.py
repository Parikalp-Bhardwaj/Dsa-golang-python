expression = "2-1-1"

def diffWaysToCompute(expression):
    if expression.isnumeric():
        return [int(expression)]
        
    output = []
    for i in range(len(expression)):
        if expression[i] == "+" or expression[i] == "-" or expression[i] == "*":
            left=diffWaysToCompute(expression[:i])
            right = diffWaysToCompute(expression[i+1:])
            for x in left:
                for y in right:
                    if expression[i] == "+":
                        output.append(x  +  y)
                    elif expression[i] == "-":
                        output.append(x  -  y)
                    else:
                        output.append(x  *  y)
    
    return output


print(diffWaysToCompute(expression))