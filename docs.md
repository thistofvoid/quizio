users:

- user-groups-id

groups:

- id
- name
- visibility [public, private]

quiz

- id
- name
- description
- visibility [public-noauth, group-only, public-auth]
- quiz-password
- is-deleted

quiz-entries

- id
- quiz-id
- text
- answers [answer_id, answer_text]
- correct-answer [andswer_id, is_correct]
