## <span style="color:Skyblue; size : 20px">*Lem-in*</span>

### <span style="color:yellow">Description

- Lem-in is meant to make a code for a digital version of an ant farm.
- The program read the contents of the file supplied as a argument and every movement the ants make from room to room after it has successfully determined the fastest route.

 
### <span style="color:yellow">Authors

- Wadia Jeglaoui
- Imane Zahid
- Fatima Ezzahra Ennadafy
- Hasnae Lamrani


### <span style="color:yellow">Usage

1. Clone the repository:
``` 
git clone https://learn.zone01oujda.ma/git/wjeglaou/lem-in
```
2. Navigate to the project directory:
```
cd lem-in
```
3. Run the program :
```
go run . test_name.txt
```

### <span style="color:yellow">Implementation Details

- A room will never start with the letter L or with # and must have no spaces.
- A tunnel joins only two rooms together never more than that.
- A room can be linked to multiple rooms.
- Two rooms can't have more than one tunnel connecting them.
- Each room can only contain one ant at a time (except at ##start and ##end which can contain as many ants as necessary).
- Each tunnel can only be used once per turn.

### <span style="color:yellow">Requirements

   - The program must be written in Go.
   - Only the [standard Go](https://pkg.go.dev/std) packages are allowed.
   - The code must respect the good practices.
   - It is recommended to have test files for unit testing.
   - The program must handle errors carefully. In no way can it quit in an unexpected manner.
   - Any unknown command will be ignored.