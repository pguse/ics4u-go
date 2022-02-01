# Working with Numbers

## Exercises

## 01-0: Student Average

1.  In Visual Studio Code, create a folder called **StudentAverage** and then run the command dotnet new console in the terminal/console window. 
2.  Modify the **Main()** function with the following starter code

```csharp
static void Main(string[] args)
{
    var mark1 = 75;
    var mark2 = 85;
    var mark3 = 90;
    // Casting - Sum of Integers to Float
    var mean = (float)(mark1 + mark2 + mark3) / 3;

    Console.WriteLine("Your average is: {0:0.00}", mean);
    //Console.WriteLine(mean.GetType());
}
```

3.  Modify the program so that the user can **input** any three marks.  Use the **int.Parse()** method to convert string input to an integer.  For example:  **int.Parse( Console.ReadLine() );**
4.  Output the average with **no decimal places**.
5.  Output all three marks and the average (all with no decimal places) using one Console.WriteLine() method.  Make sure you include words that describe what the numbers represent.

## 01-1: Area of a Circle
In Visual Studio Code, create a folder called **CircleArea** and then run the command dotnet new console in the terminal/console window.  The user should enter the **radius** of a circle, and the program should output the circle's **area**.  **Note:** Use **Math.PI** to represent the value of pi and **Math.Pow()** to perform the exponentiation.

### For example :
```
Radius:  5
Area:  78.53981633974483
```

## 01-2: Hypotenuse of a Right-Angled Triangle
In Visual Studio Code, create a folder called **Hypotenuse** and then run the command dotnet new console in the terminal/console window.  The user should enter the base and height of a right angled-triangle, and the program should output the length of the triangle's hypotenuse.  **Note:** Use **Math.Sqrt()** to do the square root calculation.

### For example :
```
Base:  3
Height:  4
Hypotenuse: 5.0
```

## 01-3: Swap Digits *(using the modulus % operator)*
In Visual Studio Code, create a folder called **Swap** and then run the command dotnet new console in the terminal/console window.  The user should enter a **two-digit** number (example: 17), and the program should output the **number** with **swapped digits** (example: 71).  

### For example :
```
Number:  17
Swapped Digits:  71
