#!/usr/bin/env bash
if [[ -n "$DRONE_TAG" ]]; then
  echo -e "\nsonar.projectVersion=$DRONE_TAG" >> sonar-project.properties
else
  echo -e "\nsonar.projectVersion=latest" >> sonar-project.properties
fi