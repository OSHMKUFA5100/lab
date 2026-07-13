```json
{
  "$schema": "https://unpkg.com/oh-my-opencode-slim@latest/oh-my-opencode-slim.schema.json",
  "preset": "zhipu",
  "presets": {
    "opencode-go": {
      "orchestrator": {
        "model": "opencode-go/glm-5.2",
        "skills": [
          "*"
        ],
        "mcps": [
          "*",
          "!context7"
        ]
      },
      "oracle": {
        "model": "deepseek/deepseek-v4-pro",
        "variant": "max",
        "skills": [
          "simplify"
        ],
        "mcps": []
      },
      "council": {
        "model": "opencode-go/qwen3.7-max",
        "variant": "high",
        "skills": [],
        "mcps": []
      },
      "librarian": {
        "model": "opencode-go/minimax-m3",
        "skills": [],
        "mcps": [
          "websearch",
          "context7",
          "gh_grep"
        ]
      },
      "explorer": {
        "model": "opencode-go/mimo-v2.5-pro",
        "skills": [],
        "mcps": []
      },
      "designer": {
        "model": "opencode-go/kimi-k2.7-code",
        "variant": "medium",
        "skills": [],
        "mcps": []
      },
      "fixer": {
        "model": "opencode-go/deepseek-v4-flash",
        "variant": "high",
        "skills": [],
        "mcps": []
      },
      "observer": {
        "model": "opencode-go/kimi-k2.7-code",
        "skills": [],
        "mcps": []
      }
    },
    "zhipu": {
      "orchestrator": {
        "model": "zhipuai-coding-plan/glm-5.2",
        "skills": [
          "*"
        ],
        "mcps": [
          "*",
          "!context7"
        ]
      },
      "oracle": {
        "model": "zhipuai-coding-plan/glm-5.2",
        "variant": "max",
        "skills": [
          "simplify"
        ],
        "mcps": []
      },
      "council": {
        "model": "zhipuai-coding-plan/glm-5.1",
        "skills": [],
        "mcps": []
      },
      "librarian": {
        "model": "zhipuai-coding-plan/glm-5-turbo",
        "skills": [],
        "mcps": [
          "websearch",
          "context7",
          "gh_grep"
        ]
      },
      "explorer": {
        "model": "zhipuai-coding-plan/glm-5-turbo",
        "skills": [],
        "mcps": []
      },
      "designer": {
        "model": "zhipuai-coding-plan/glm-5.2",
        "variant": "medium",
        "skills": [],
        "mcps": []
      },
      "fixer": {
        "model": "zhipuai-coding-plan/glm-5.2",
        "skills": [],
        "mcps": []
      },
      "observer": {
        "model": "zhipuai-coding-plan/glm-4.6v",
        "skills": [],
        "mcps": []
      }
    },
    "deepseek": {
      "orchestrator": {
        "model": "deepseek/deepseek-v4-pro",
        "variant": "high",
        "skills": [
          "*"
        ],
        "mcps": [
          "*",
          "!context7"
        ]
      },
      "oracle": {
        "model": "deepseek/deepseek-v4-pro",
        "variant": "max",
        "skills": [
          "simplify"
        ],
        "mcps": []
      },
      "council": {
        "model": "deepseek/deepseek-v4-pro",
        "variant": "medium",
        "skills": [],
        "mcps": []
      },
      "librarian": {
        "model": "deepseek/deepseek-v4-flash",
        "skills": [],
        "mcps": [
          "websearch",
          "context7",
          "gh_grep"
        ]
      },
      "explorer": {
        "model": "deepseek/deepseek-v4-flash",
        "skills": [],
        "mcps": []
      },
      "designer": {
        "model": "deepseek/deepseek-v4-pro",
        "variant": "medium",
        "skills": [],
        "mcps": []
      },
      "fixer": {
        "model": "deepseek/deepseek-v4-flash",
        "variant": "high",
        "skills": [],
        "mcps": []
      }
    }
  },
  "disabled_agents": []
}

```

