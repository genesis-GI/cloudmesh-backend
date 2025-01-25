from ollama import chat

prompt = "Hello, how are you?"

stream = chat(
    model='genesis',
    messages=[{
      'role': 'user', 
      'content': prompt
    }],
    stream=True,
)

for chunk in stream:
  print(chunk['message']['content'], end='', flush=True)