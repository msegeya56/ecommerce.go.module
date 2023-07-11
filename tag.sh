#!/bin/bash

# Get the current version from the module
version=$(git describe --abbrev=0 --tags)

# Increment the version number
IFS='.' read -r -a version_parts <<< "$version"
major="${version_parts[0]}"
minor="${version_parts[1]}"
patch="${version_parts[2]}"
minor=$((minor + 1))
new_version="$major.$minor.$patch"

# Create the new tag
git tag "$new_version"

# Push the new tag to the repository
git push origin "$new_version"

# Push the updates to the 'main' branch
git push origin main
