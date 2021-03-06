package fixtures

import "fmt"

func RunRSyncScript() string {
	return fmt.Sprintf(runnerScript, forward, "\n\trsync -a /tmp/local/ /home/vcap/app/", rsyncRunningToLocal)
}

func CommitScript() string {
	return fmt.Sprintf(runnerScript, "", "", "")
}

func StageRSyncScript() string {
	return fmt.Sprintf(stageScript, "", "\n\trsync -a /tmp/app/ /tmp/local/")
}

const forward = `
	echo 'Forwarding: some-name some-other-name'
	sshpass -p 'some-code' ssh -f -N \
		-o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no \
		-o LogLevel=ERROR -o ExitOnForwardFailure=yes \
		-o ServerAliveInterval=10 -o ServerAliveCountMax=60 \
		-p 'some-port' 'some-user@some-ssh-host' \
		-L 'some-from:some-to' \
		-L 'some-other-from:some-other-to'`

const rsyncRunningToLocal = `
	if [[ -z $(ls -A /tmp/local) ]]; then
		rsync -a /home/vcap/app/ /tmp/local/
	fi`

const runnerScript = `
	set -e%s%s
	if [[ ! -z $(ls -A /home/vcap/app) ]]; then
		exclude='--exclude=./app'
	fi
	tar $exclude -C /home/vcap -xzf /tmp/droplet
	chown -R vcap:vcap /home/vcap%s
	command=$1
	if [[ -z $command ]]; then
		command=$(jq -r .start_command /home/vcap/staging_info.yml)
	fi
	exec /tmp/lifecycle/launcher /home/vcap/app "$command" ''
`

const stageScript = `
	set -e
	chown -R vcap:vcap /tmp/app /tmp/cache
	%ssu vcap -p -c "PATH=$PATH exec /tmp/lifecycle/builder -buildpackOrder $0 -skipDetect=$1"%s
`
