# world-wizard

A fun side project born out of the question: "Can you hack Wordle ?". Whilst this is not a hack, it does help solve Wordle puzzles, by generating possible answers given the current board state as input. 

# Usage

The Wordle board state is the input. This is translated to the following command line arguments  
**-e** characters excluded (dark grey blocks), these can be entered as a sequence with no seperator e.g: abcd  
**-i** characters included but whose position is unknown (yellow blocks), these can be entered as a sequence with no seperator e.g: abcd  
**-p** characters included whose positions are known (green blocks), these can be entered as a sequence consisting of the characters followed by the index e.g: a0d3  

# Example
Using STEAK as a starting point, the Wordle board now looks like:

![Wordle_1](https://user-images.githubusercontent.com/38185025/224811464-62b61cce-42d0-412f-8f2f-8e95fc42d341.jpeg)

From the result, it is seen that 's' is in position 0 and the corrector word does not contain 't', 'e', 'a' or 'k', this can be translated to the command: `.\wordleWizard.exe -p s0 -e teak` which outputs:
![image](https://user-images.githubusercontent.com/38185025/224812354-2c81d2ab-b279-4e92-99ef-40c26c941865.png)

One of the retrieved words could then be used as a next guess. Assuming SHOWN is used as the next input, leaving the state of the board as follows:  
![Wordle_2](https://user-images.githubusercontent.com/38185025/224812888-4f2b6bf4-87a7-4274-bd7c-9a6b87455544.jpeg)

From the last guess, the correct position of 'o' is revealed as index 2, additionally it is known that 'h', 'w' and 'n' are also excluded, this can be used to craft a more precise command: `.\wordleWizard.exe -p s0o2 -e teakhwn` which outputs:
![image](https://user-images.githubusercontent.com/38185025/224813451-3d03f9b5-1ae7-4ad3-baa3-bfd0e0240aaa.png)  
As more info is added, the number of possibilities reduce. 

Using SPOIL as the next guess, results in the following board:
![Wordle_3](https://user-images.githubusercontent.com/38185025/224814167-b492a6e1-52d4-4762-b279-6d18b93f8570.jpeg)

Incorporating the new info of 2 more exclusions 'p' and 'i' as well as an inclusion 'l' whose position is not known, an updated query can be run: `.\wordleWizard.exe -p s0o2 -e teakhwnpi -i l`
![image](https://user-images.githubusercontent.com/38185025/224814716-42a36ff7-1aa6-44d8-a944-c7e1cb4524e3.png)

![Wordle_4](https://user-images.githubusercontent.com/38185025/224815149-9eb5f76b-1165-459e-b7dd-41004459a37e.jpeg)

In this case, the guessed word works!
