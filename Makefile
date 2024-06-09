codegen:
	npx graphql-codegen-esm -w

rds-start:
	aws rds start-db-instance --db-instance-identifier testing

rds-stop:
	aws rds stop-db-instance --db-instance-identifier testing

integration-test:
	cd server/test && go run .

development:
	tmux \; \
		new-session -d -s dev-session 'cd server && air' \; \
		split-window -h 'cd client && npm run development' \; \
		split-window -v 'npx graphql-codegen-esm -w' \; \
		select-pane -t 0 \; \
		attach-session -t dev-session

qa:
	tmux \; \
		new-session -d -s qa-session 'cd server && go run .' \; \
		split-window -h 'cd client && npm run build && npm run start' \; \
		select-pane -t 0 \; \
		attach-session -t qa-session