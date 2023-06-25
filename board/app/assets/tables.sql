CREATE TABLE IF NOT EXISTS boards (
	id UUID UNIQUE  NOT NULL ,
	workspace_id UUID  NOT NULL,
	title VARCHAR NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP DEFAULT  NULL,
	deleted_at TIMESTAMP DEFAULT  NULL,
	CONSTRAINT boards_pkey PRIMARY KEY (id)
); 
CREATE TABLE IF NOT EXISTS sections (
	id UUID UNIQUE  NOT NULL,
	board_id UUID  NOT NULL,
	title VARCHAR NOT NULL,
	sort_order SMALLINT NOT NULL CHECK (sort_order >= 0),
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP DEFAULT  NULL,
	deleted_at TIMESTAMP DEFAULT  NULL,
	CONSTRAINT sections_pkey PRIMARY KEY (id)
	
);
CREATE TABLE IF NOT EXISTS tasks (
	id UUID UNIQUE  NOT NULL,
	title VARCHAR NOT NULL,
	section_id UUID  NOT NULL,
	sort_order SMALLINT NOT NULL CHECK (sort_order >= 0),
	creator_id UUID  NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP DEFAULT NULL,
	deleted_at TIMESTAMP DEFAULT  NULL,
	CONSTRAINT tasks_pkey PRIMARY KEY (id)
	
);
CREATE TABLE IF NOT EXISTS sub_tasks (
	id UUID UNIQUE  NOT NULL,
	task_id UUID  NOT NULL,
	title VARCHAR NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP DEFAULT  NULL,
	deleted_at TIMESTAMP DEFAULT  NULL,
	CONSTRAINT sub_tasks_pkey PRIMARY KEY (id)
	
);
CREATE TABLE IF NOT EXISTS tags (
	id UUID UNIQUE  NOT NULL,
	title VARCHAR NOT NULL,
	pattern VARCHAR DEFAULT  NULL,
	sort_order SMALLINT NOT NULL CHECK (sort_order >= 0),
	board_id UUID  NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP DEFAULT  NULL,
	deleted_at TIMESTAMP DEFAULT  NULL,
	CONSTRAINT tag_pkey PRIMARY KEY (id)
	
);


ALTER TABLE "sections" ADD CONSTRAINT fk_board_id FOREIGN KEY(board_id) REFERENCES boards(id);
ALTER TABLE "tasks" ADD CONSTRAINT fk_section_id FOREIGN KEY(section_id) REFERENCES sections(id);

ALTER TABLE "sub_tasks" ADD CONSTRAINT fk_task_id FOREIGN KEY(task_id) REFERENCES tasks(id);

ALTER TABLE "tags" ADD CONSTRAINT fk_board_id FOREIGN KEY(board_id) REFERENCES boards(id);
