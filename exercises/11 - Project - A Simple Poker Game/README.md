# A Simple Poker Game

This is an example of a **Go** program that has been organized using multiple files.

##
Here is your task:

Create a **folder** on your computer called **poker**. Copy the following **Go** files:  **main.go**, **card.go**, **hand.go**, and **deck.go**; from here into your **poker** folder.  In **Visual Studio Code** open your **poker** folder, then create a **New Terminal**.

Your **poker** folder should now look like this:

![poker folder](https://github.com/pguse/ics4u-go/blob/main/exercises/11%20-%20Project%20-%20A%20Simple%20Poker%20Game/pokerFolderTerminal.png)

In order to run our program, we need to do two new steps.  First, we will define our project as a **module**.  We do this in the **Terminal** using the command,

```
go mod init poker
```

Now your **poker** folder should look like this

![poker folder with go.mod file](https://github.com/pguse/ics4u-go/blob/main/exercises/11%20-%20Project%20-%20A%20Simple%20Poker%20Game/pokerFolderModuleTerminal.png)

The **go.mod** file that you have created looks like this when you open it.

![go mod file](https://github.com/pguse/ics4u-go/blob/main/exercises/11%20-%20Project%20-%20A%20Simple%20Poker%20Game/pokerFolderModuleFile.png)

It just contains the name of the module **poker** and the version of **Go** you were using when creating the file.

The second new step involves how we run our project.  In the **Terminal**, when we are within the **poker** folder, type

```
go run .
```

and you should see the following output.

![go run file](https://github.com/pguse/ics4u-go/blob/main/exercises/11%20-%20Project%20-%20A%20Simple%20Poker%20Game/pokerFolderRun.png)


## Exercises

Complete the following exercises.  When completed, copy your **poker** folder to the **14 - A Simple Poker Game** shared folder on OneDrive.

### 11-0: Flush-Straight-StraightFlush-Pair Functions

Go back to the **Struct** exercises and copy your code from your **Flush**, **Straight**, **StraightFlush**, and **Pair** methods into the **hand.go** file.

### 11-1: Three-of-A-Kind

Create a **ThreeOfAKind** method in the **hand.go** file.

```go
func (h Hand) ThreeOfAKind() bool {
	// Three cards with the same rank - assume a sorted hand
	return false
}
```

### 11-2: Four-of-A-Kind

Create a **FourOfAKind** method in the **hand.go** file.

```go
func (h Hand) FourOfAKind() bool {
	// Four cards with the same rank - assume a sorted hand
	return false
}
```

### 11-3: Full House

Create a **FullHouse** method in the **hand.go** file.

```go
func (h Hand) FullHouse() bool {
	// Two cards with the same rank and Three cards with same rank - assume a sorted hand
	return false
}
```

### 11-4: Testing

Copy the **test.go** file into your **poker** folder.  Complete tests in the **test** function for the following methods:  **Straight**, **StraightFlush**, **Pair**, **ThreeOfAKind**, **FourOfAKind**, and **FullHouse**.  Run the **test** function in **main**.  Here is the code already provided in **test.go** that tests your **Flush** method.

```go
func test() {
	// A hand with nothing
	h1 := Hand{Card{"A", '♢'}, Card{"3", '♡'}, Card{"7", '♣'}, Card{"J", '♠'}, Card{"K", '♡'}}
	// A hand with a flush
	h2 := Hand{Card{"A", '♣'}, Card{"3", '♣'}, Card{"5", '♣'}, Card{"7", '♣'}, Card{"9", '♣'}}

	// Testing Flush()
	if !h1.Flush() && h2.Flush() {
		fmt.Println("Flush: OK")
	} else {
		fmt.Println("Flush: Not OK")
	}
}
```

Here is what the output should look like before editing the **test.go** file

![test output](https://github.com/pguse/ics4u-go/blob/main/exercises/11%20-%20Project%20-%20A%20Simple%20Poker%20Game/test.png)
