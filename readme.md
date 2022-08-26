### Config for test
    "go.toolsManagement.autoUpdate": true,
    "go.coverOnSave": true,
    "go.coverOnSingleTest": true,
    "go.coverageDecorator": {
        "type": "gutter",
        "coveredGutterStyle": "blockgreen",
        "uncoveredGutterStyle": "blockred"
    },

### Config for integration test
    "gopls": {
        "build.buildFlags": [
            "-tags=integration"
        ]
    },