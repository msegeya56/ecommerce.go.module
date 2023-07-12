!/bin/sh

# Update go mod

go mod tidy

# Get the most recent tag
latest_tag=$(git describe --abbrev=0 --tags)

# Increment the tag version
new_tag=$(echo $latest_tag | awk -F. -v OFS=. '{$NF++;print}')

# Add the files you want to commit

git add .

# Commit the changes with a descriptive message
commit_message="Increment tag to $new_tag"
git commit -m "$commit_message"

# Create the new tag and push it
git tag $new_tag
git push origin $new_tag
