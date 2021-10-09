from github import Github
g = Github("ArielSixwings","")
user = g.get_user()
repo = g.get_repo("ArielSixwings/GoDeep")

all_files = []

contents = repo.get_contents("")
# for i in contents:
# 	print(i)
while contents:
    file_content = contents.pop(0)
    if file_content.type == "dir":
        contents.extend(repo.get_contents(file_content.path))
    else:
        file = file_content
        all_files.append(str(file).replace('ContentFile(path="','').replace('")',''))

    print(all_files[-1])
with open('General_Data.csv', 'r') as file:
    content = file.read()

# Upload to github
git_prefix = 'folder3/'
git_file = git_prefix + 'General_Data.csv'

# contents = repo.get_contents('folder2/General_Data.csv')
# repo.update_file(contents.path, "committing files", content, contents.sha, branch="master")

if git_file in all_files:
    print("at update")
    contents = repo.get_contents(git_file)
    repo.update_file(contents.path, "committing files", content, contents.sha, branch="main")
    print(git_file + ' UPDATED')
else:
    repo.create_file(git_file, "committing files", content, branch="main")
    print(git_file + ' CREATED')