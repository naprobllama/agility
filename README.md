# agility

## Welcome to the SWAPI project I've put together for Agility.

Here are the Requirements.

The application should:
Take a Star Wars character name, and returns the following information about the
- character:
- Starship
- Starship name, cargo capacity, and Starship class
- Home Planet
- Planet name, population, and climate
- Species
- Name, language, and average lifespan

Handle as few as one letter in the input
e.g. if the user inputs c, return information for all characters with a c in the name
If multiple characters are found, return in alphabetical order by character name

Do Not:
Make use of any of the SWAPI helper libraries
You are able to use whichever language you are most comfortable with. How you do
this, how you take the input, or display the output is all up to you. Creativity is
encouraged

BE Take Home Technical Interview 2
Additional Considerations
What do you do if you get back multiple characters? Sorts the array. 
What do you do if you get no characters back? Prints farewell string. 
How to handle multiple starships? Prints them all. 
How to handle if a section (starship, home planet, species) is empty? Just ignores it. 

## Liberties I took:

There were ALOT of time constraints on completing this project.
In no way does this project reflect code I would consider production ready nor architecturally sound.
Things I would have wanted to see in the project include:
- Unit tests for the win 
- Way way more consistency on naming of variables and data. The data structures in general needed more TLC. Interfaces for extensibilty etc. 
- Error handling through out the application including reporting failed go routiunes.
- All numbers (such as channel lengths) would be dynamic
- Would have revisited using the name as the unique foreign key on child structs. Yuck. I would have preferred to see url be used.
- Remove garbage retries for my bad internet connection
- Architecture wasn't awesome. Would love to see REST URLS and HTTP get() function separated out

As with most projects, it did not get to the glistening beautiful state I wanted to see it in.
However, it works to the specs and considering that I was given the project less than 5 days ago it is nonetheless an MVP. 

I hope to have time to test this project further, very next steps would be swapping person name for URL as the foreign key.
Then a massive data naming clean up.  
It does pass my basic testing but I have not yet stress tested it. 

Overall, regardles of the state its in, the project shows my "approach" to some things. 

Thank you for taking the time to look this project over. 






