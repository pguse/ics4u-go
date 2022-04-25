# Chess Ranking Project

## Task

Mr. Profit has asked us to create a **Go** program that helps him keep track of chess player ratings in his Fall Arts Period.  The initial version of the program is provided here.  It loads student information from a file and stores the modified information in another file.

Mr. Profit would like us to improve our program in the following ways.


1. He would like us to create a menu with the following options for him to choose,
   - Update a game result
   - Add a student to the group
     - If there are an odd number of students, add Mr. Profit as a student.  Assume he has the lowest rank and that he is in Deroche.
   - Remove a student from the group
     - If there are an odd number of students, add Mr. Profit as a student.  Assume he has the lowest rank and that he is in Deroche.
   - Quit the program and store the updated results in the output.txt file.  Students should be sorted by rank, before writing to the file.
                
       - There should be a **menu** function in your program.

2. At the moment the ratings of players are swapped if it benefits the player that wins.  He would like us to modify the ratings in the following way.
..* For each win, add the opponent's rating (rank) plus 400 to the winning student's rating (rank),
..* For each loss, add the opponent's rating (rank) minus 400 to the losing student's rating (rank).