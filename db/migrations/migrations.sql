CREATE DATABASE IF NOT EXISTS cico_learn;

CREATE TYPE gender AS ENUM ('men', 'women', 'unknown');
CREATE TYPE english_level AS ENUM ('Pre-A1(foundation)', 'A1(Elementary)', 'A2(Pre-Intermediate)', 'B1(Intermediate)', 'B2(Upper (Intermediate)', 'C1(Advanced)', 'C2(Proficiency)');
CREATE TYPE language AS ENUM ('indonesia', 'english');
CREATE TYPE target AS ENUM ('Learning for fun', 'Job interview', 'Study abroad', 'Fluent communication', 'Understand university lecture', 'Work');
CREATE TYPE time_target AS ENUM ('5 minutes', '10 minutes', '15 minutes', '30 minutes', '45 minutes', '1 hour', '2 hours');
-- USERS TABLE
CREATE TABLE IF NOT EXISTS users(
    id SERIAL PRIMARY KEY,
    first_name VARCHAR NOT NULL,
    last_name VARCHAR,
    username VARCHAR UNIQUE NOT NULL,
    email VARCHAR UNIQUE NOT NULL,
    address VARCHAR,
    gender gender NOT NULL DEFAULT 'unknown',
    profile_image VARCHAR, 
    background_image VARCHAR, 
    level english_level DEFAULT 'Pre-A1(foundation)',
    langauge language DEFAULT 'english' NOT NULL,
    target target DEFAULT 'Learning for fun',
    daily_reminder time with time zone[5],
    time_target time_target DEFAULT '5 minutes',
    password_hash VARCHAR NOT NULL,
    updated_at timestamp with time zone DEFAULT now(), 
    create_at timestamp with time zone DEFAULT now(), 
);

CREATE TYPE vocabulary_type AS ENUM ('word', 'expression');
-- VOCABULARY TABLE
CREATE TABLE IF NOT EXISTS vocabulary(
    id SERIAL PRIMARY KEY,
    topic VARCHAR,
    description VARCHAR,
    type vocabulary_type NOT NULL,
)

-- WORDS TABLE
CREATE TABLE IF NOT EXISTS words(
    id SERIAL PRIMARY KEY,
    image VARCHAR,
    word VARCHAR,
    meaning VARCHAR,
    synonyms VARCHAR,
    sentence VARCHAR,
    -- vocabulary_id INT REFERENCES vocabulary(id) ON DELETE SET NULL
)

CREATE TABLE IF NOT EXISTS vocabulary_word(
    id SERIAL PRIMARY KEY,
    vocabulary_id INT REFERENCES vocabulary(id) ON DELETE SET NULL,
    word_id INT REFERENCES words(id) ON DELETE CASCADE
)

CREATE TYPE expression_type AS ENUM ('phrasal verbs', 'collocations', 'phrases', 'idoms', 'sayings', 'proverbs')
-- EXPRESSIONS TABLE
CREATE TABLE IF NOT EXISTS expressions(
    id SERIAL PRIMARY KEY,
    phrase VARCHAR,
    meaning VARCHAR
    sentence VARCHAR
    type expression_type NOT NULL,
    -- vocabulary_id INT REFERENCES vocabulary(id) ON DELETE SET NULL
)

CREATE TABLE IF NOT EXISTS vocabulary_expression(
    id SERIAL PRIMARY KEY,
    vocabulary_id INT REFERENCES vocabulary(id) ON DELETE SET NULL
    expression_id INT REFERENCES expressions(id) ON DELETE SET NULL
)

-- WORD EXAMPLES TABLE
CREATE TABLE IF NOT EXISTS word_examples(
    id SERIAL PRIMARY KEY,
    example VARCHAR,
    meaning VARCHAR,
    word_id INT REFERENCES words(id) ON DELETE CASCADE
)

-- EXPRESSION EXAMPLES TABLE
CREATE TABLE IF NOT EXISTS expression_examples(
    id SERIAL PRIMARY KEY,
    example VARCHAR,
    meaning VARCHAR,
    expression_id INT REFERENCES expressions(id) ON DELETE CASCADE,
)

CREATE TABLE IF NOT EXISTS user_type_words(
    id SERIAL PRIMARY KEY,
    correct boolean
    user_id INT REFERENCES users(id) ON DELETE NO ACTION,
    word_id INT REFERENCES words(id) ON DELETE NO ACTION,
    created_at timestamp with time zone DEFAULT now()
)

CREATE TABLE IF NOT EXISTS user_type_expressions(
    id SERIAL PRIMARY KEY,
    correct boolean
    user_id INT REFERENCES users(id) ON DELETE NO ACTION,
    expression_id INT REFERENCES expressions(id) ON DELETE NO ACTION,
    created_at timestamp with time zone DEFAULT now()
)

CREATE TABLE IF NOT EXISTS user_type_three_sentences_contain_word(
    id SERIAL PRIMARY KEY,
    first_sentence VARCHAR,
    second_sentence VARCHAR,
    third_sentence VARCHAR,
    feedback VARCHAR,
    score DECIMAL(5, 2),
    word_id INT REFERENCES words(id) ON DELETE NO ACTION,
    user_id INT REFERENCES users(id) ON DELETE NO ACTION,
    created_at timestamp with time zone DEFAULT now()
)

CREATE TABLE IF NOT EXISTS user_type_three_sentences_contain_expression(
    id SERIAL PRIMARY KEY,
    first_sentence VARCHAR,
    second_sentence VARCHAR,
    third_sentence VARCHAR,
    feedback VARCHAR,
    score DECIMAL(5, 2),
    expression_id INT REFERENCES expressions(id) ON DELETE NO ACTION,
    user_id INT REFERENCES users(id) ON DELETE NO ACTION,
    created_at timestamp with time zone DEFAULT now()
)

CREATE TABLE IF NOT EXISTS user_speak_sentence_contain_expression(
    id SERIAL PRIMARY KEY,
    accuracy DECIMAL(5, 2),
    feedback VARCHAR,
    expression_id INT REFERENCES expressions(id) ON DELETE NO ACTION,
    user_id INT REFERENCES users(id) ON DELETE NO ACTION,
    created_at timestamp with time zone DEFAULT now()
)

CREATE TABLE IF NOT EXISTS user_speak_sentence_contain_word(
    id SERIAL PRIMARY KEY,
    accuracy DECIMAL(5, 2),
    feedback VARCHAR,
    word_id INT REFERENCES words(id) ON DELETE NO ACTION,
    user_id INT REFERENCES users(id) ON DELETE NO ACTION,
    created_at timestamp with time zone DEFAULT now()
)

-- DIALOGUES TABLE
CREATE TABLE IF NOT EXISTS dialogues(
    id SERIAL PRIMARY KEY,
    topic VARCHAR,
    description VARCHAR,
)

CREATE TABLE IF NOT EXISTS dialog_items(
    id SERIAL PRIMARY KEY,
    order SMALLINT,
    dialog VARCHAR,
    dialog_id INT REFERENCES dialogues(id) ON DELETE CASCADE
)

CREATE TABLE IF NOT EXISTS user_learn_dialog(
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE NO ACTION,
    dialog_id INT REFERENCES dialogues(id) ON DELETE CASCADE
    -- overall_score DECIMAL(5, 2)
)

CREATE TABLE IF NOT EXISTS user_speak_dialog(
    id SERIAL PRIMARY KEY,
    accuracy DECIMAL(5, 2),
    feedback VARCHAR,
    dialog_item_id INT REFERENCES dialog_items(id) ON DELETE NOT ACTION
    user_id INT REFERENCES users(id) ON DELETE NO ACTION,
)

-- GRAMMAR LESSONS TABLE
CREATE TABLE IF NOT EXISTS grammar_lessons(
    id SERIAL PRIMARY KEY,
    topic VARCHAR,
    description VARCHAR,
)

CREATE TABLE IF NOT EXISTS grammar_sections(
    id SERIAL PRIMARY KEY,
    section_order SMALLINT,
    title VARCHAR,
    grammar_lesson_id INT REFERENCES grammar_lessons(id) ON DELETE CASCADE
)

CREATE TYPE grammar_section_content_type AS ENUM ('header', 'text')
CREATE TABLE IF NOT EXISTS grammar_section_explanations(
    id SERIAL PRIMARY KEY,
    order SMALLINT,
    content VARCHAR,
    type grammar_section_content_type,
    grammar_section_id INT REFERENCES grammar_sections(id) ON DELETE CASCADE
)

CREATE TYPE correct_answer_choice ('a_choice', 'b_choice', 'c_choice', 'd_choice',)
CREATE TABLE IF NOT EXISTS grammar_exercises(
    id SERIAL PRIMARY KEY,
    order SMALLINT,
    question VARCHAR,
    a_choice VARCHAR,
    b_choice VARCHAR,
    c_choice VARCHAR,
    d_choice VARCHAR,
    correct_answer correct_answer_choice,
    grammar_section_id INT REFERENCES grammar_sections(id) ON DELETE CASCADE,
    grammar_lesson_id INT REFERENCES grammar_lessons(id) ON DELETE CASCADE
    create_at timestamp with time zone DEFAULT now(), 
)

CREATE TABLE IF NOT EXISTS user_learn_grammar_lessons(
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE NO ACTION,
    grammar_lesson_id INT REFERENCES grammar_lessons(id) ON DELETE RESTRICT,
)

CREATE TABLE IF NOT EXISTS user_do_exercises(
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE NO ACTION,
    grammar_exercise_id INT REFERENCES grammar_exercises(id) ON DELETE NO ACTION,
)

CREATE TABLE IF NOT EXISTS user_open_grammar_sections(
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE NO ACTION,
    grammar_section_id INT REFERENCES grammar_sections(id) ON DELETE NO ACTION,
)

-- NOTIFICATIONS TABLE
CREATE TABLE IF NOT EXISTS notifications(
    id SERIAL PRIMARY KEY,
    title VARCHAR,
    emoticon VARCHAR,
    message VARCHAR,
    send_at timestamp with time zone,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
)

CREATE TABLE IF NOT EXISTS user_get_notifications(
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE NO ACTION,
    notification_id INT REFERENCES notifications(id) ON DELETE CASCADE
)


CREATE TYPE train_with_ai_mode AS ENUM ('free talk', 'challenge', 'role play')

-- TRAIN WITH AI TOPICS TABLE
CREATE TABLE IF NOT EXISTS train_with_ai_topics(
    id SERIAL PRIMARY KEY,
    image VARCHAR,
    scenario_name VARCHAR,
    scenario VARCHAR,
    mode train_with_ai_mode NOT NULL,
)

-- TRAIN WITH AI DIALOGUES TABLE
CREATE TABLE IF NOT EXISTS user_train_with_ai_dialogues(
    id SERIAL PRIMARY KEY,
    mode train_with_ai_mode NOT NULL,
    user_id INT REFERENCES users(id) ON DELETE NO ACTION,
    train_with_ai_topic_id INT REFERENCES train_with_ai_topics(id) ON DELETE NO ACTION,
    started_at timestamp with time zone,
    last_opened_at timestamp with time zone,
)

CREATE TABLE IF NOT EXISTS train_with_ai_dialog_items(
    id SERIAL PRIMARY KEY,
    order_dialog SMALLINT,
    dialog VARCHAR,
    is_from_user boolean,
    train_with_ai_dialog_id INT REFERENCES train_with_ai_dialogues(id) ON DELETE NO ACTION
)

-- TRAIN WITH AI RULES TABLE
CREATE TABLE IF NOT EXISTS train_with_ai_rules(
    id SERIAL PRIMARY KEY,
    rule VARCHAR,
    train_with_ai_topic_id INT REFERENCES train_with_ai_topics(id) ON DELETE CASCADE
)

CREATE TABLE IF NOT EXISTS memorize_vocabulary_drafts(
    id SERIAL PRIMARY KEY,
    name VARCHAR,
    image VARCHAR,
    user_id INT REFERENCES users(id) ON DELETE NO ACTION
)

CREATE TABLE IF NOT EXISTS words_in_draft(
    id SERIAL PRIMARY KEY,
    memorize_vocabulary_draft_id INT REFERENCES memorize_vocabulary_drafts(id) ON DELETE CASCADE,
    word_id INT REFERENCES words(id) ON DELETE RESTRICT,
    added_at timestamp with time zone DEFAULT now()
)

CREATE TABLE IF NOT EXISTS expressions_in_draft(
    id SERIAL PRIMARY KEY,
    memorize_vocabulary_draft_id INT REFERENCES memorize_vocabulary_drafts(id) ON DELETE CASCADE,
    expression_id INT REFERENCES expressions(id) ON DELETE RESTRICT,
    added_at timestamp with time zone DEFAULT now(),
)

CREATE TABLE IF NOT EXISTS user_memorize_word_type(
    id SERIAL PRIMARY KEY,
    is_correct boolean,
    word_id INT REFERENCES words(id) ON DELETE NO ACTION,
    memorize_vocabulary_draft_id INT REFERENCES memorize_vocabulary_drafts(id) ON DELETE NO ACTION,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
)

CREATE TABLE IF NOT EXISTS user_memorize_expression_type(
    id SERIAL PRIMARY KEY,
    is_correct boolean,
    expression_id INT REFERENCES expressions(id) ON DELETE NO ACTION,
    memorize_vocabulary_draft_id INT REFERENCES memorize_vocabulary_drafts(id) ON DELETE NO ACTION,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
)

CREATE TABLE IF NOT EXISTS user_memorize_word_type_two_sentences(
    id SERIAL PRIMARY KEY,
    score DECIMAL(5, 2),
    feedback VARCHAR,
    word_id INT REFERENCES words(id) ON DELETE NO ACTION,
    memorize_vocabulary_draft_id INT REFERENCES memorize_vocabulary_drafts(id) ON DELETE NO ACTION,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
)

CREATE TABLE IF NOT EXISTS user_memorize_expression_type_two_sentences(
    id SERIAL PRIMARY KEY,
    score DECIMAL(5, 2),
    feedback VARCHAR,
    expression_id INT REFERENCES expressions(id) ON DELETE NO ACTION,
    memorize_vocabulary_draft_id INT REFERENCES memorize_vocabulary_drafts(id) ON DELETE NO ACTION,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
)

CREATE TABLE IF NOT EXISTS user_memorize_word_speak_sentence(
    id SERIAL PRIMARY KEY,
    accuray DECIMAL(5, 2),
    feedback VARCHAR,
    word_id INT REFERENCES words(id) ON DELETE NO ACTION,
    memorize_vocabulary_draft_id INT REFERENCES memorize_vocabulary_drafts(id) ON DELETE NO ACTION,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
)

CREATE TABLE IF NOT EXISTS user_memorize_expression_speak_sentence(
    id SERIAL PRIMARY KEY,
    accuray DECIMAL(5, 2),
    feedback VARCHAR,
    expression_id INT REFERENCES expressions(id) ON DELETE NO ACTION,
    memorize_vocabulary_draft_id INT REFERENCES memorize_vocabulary_drafts(id) ON DELETE NO ACTION,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
)