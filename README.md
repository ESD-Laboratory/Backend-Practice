# Backend-Practice
Backend Repository for Practice 

## Initial Setup
Follow the steps below to setup the project on your local machine.
- Fork this repository to your account
- Clone the forked repository to your local machine
- Create a new branch before making any changes to the code

## Assignments
All assignments will be pushed into the Assignments folder, each with its own subfolder. Follow the instructions in each assignment folder

## Submission

- Before you do the assignment first you need to create a new branch with your name and the assignment name e.g. `git checkout -b <your-name>/<assignment-name>`
- After you are done with the assignment, push the branch to your forked repository
- Create a pull request from your forked repository to the original repository

## Syncing your forked repository
Original repository will be periodically updated and you will need to manually sync new changes from your forked, private repository.

In your local machine, set the upstream to original repository
```bash
cd <your-local-repository>
git remote add upstream https://github.com/ESD-Laboratory/Backend-Practice
```
then checkout to `main` branch and pull the changes from the upstream
```bash
git checkout main
git pull upstream main
```
then push the changes to your forked repository
```bash
git push origin main
```

## Resources
- [Git and GitHub](https://www.youtube.com/watch?v=SWYqp7iY_Tc)
- [Go](https://golang.org/doc/)
- [Go by Example](https://gobyexample.com/)


### Sincerly,
ESD Laboratory Software Development Team