package app

import agentService "agent/internal/agent"

func Run() {
	agentService.StartAgent()
}