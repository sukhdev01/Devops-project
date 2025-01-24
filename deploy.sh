#!/bin/bash
set -eou pipefail

NAMESPACE="devops-project"
TIMEOUT="300s"
CHARTNAME="devops-project"
FILENAME="values.yaml"

WORK_DIR="Documents/Devops-project"

echo -e "Running Heva..."
cd $HOME/$WORK_DIR/charts/devops-project 


echo -e "Rendering helm templates based on supplied values file..."
helm template -v 5 --namespace $NAMESPACE --logtostderr --debug --values $FILENAME . 2>&1

if [[ $? -ne 0 ]]; then
    echo "Helm Template validation failed. Hence release is halted!"
    exit 1
fi

echo -e "Doing Release..."
helm upgrade -v 5 --namespace $NAMESPACE --logtostderr --debug --install --atomic --timeout ${TIMEOUT} --cleanup-on-fail --values $FILENAME $CHARTNAME . 2>&1 
