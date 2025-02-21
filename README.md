```
• Method: GET

• GET /palindrome:

• Query Param: text

• Response:
o "Palindrome" for palindromes.
o "Not palindrome" for non-palindromes.
```

```
• GET /:
 Response: Hello Go developers
```

```
• Method: GET
• GET /language:
• Response: JSON object containing language.
```

```
• Method: GET
• GET /language/{id}:
• Response: JSON object containing language details.
```

```
• Method: POST
• POST /language/add:
• Response: JSON object containing the newly added language.
```

```
• PATCH /language/{id}:
• Request Body: JSON object with updated language data.
• Response: Updated language data in JSON format.
• Method: PATCH
```

```
• DELETE /language/{id}:
• Response: HTTP 204 No Content on successful deletion.
• Method: DELETE

```
