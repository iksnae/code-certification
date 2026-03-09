---
name: web-search
description: Perform live web searches using the Brave Search MCP tool to find current, factual, and up-to-date information from the internet.
---

# Brave Web Search Skill (MCP)

This skill enables Roo to perform real-time web searches using the **Brave Search MCP tool** when the user asks for information that cannot be reliably answered from existing knowledge alone.

## When to Use This Skill

Activate this skill when the user requests:
- Current events or recent news
- Live facts, statistics, or up-to-date documentation
- Information that may have changed after the model’s training cutoff
- General web searches such as “look this up”, “search the web”, or “find sources”

Do **not** use this skill for:
- Pure reasoning tasks
- Code generation that does not require external facts
- Questions clearly answerable from general knowledge

## Search Strategy

When using the Brave Search MCP tool:

1. Form a **concise, neutral search query**
   - Avoid unnecessary filler words
   - Prefer factual phrasing over conversational phrasing

2. Execute **one search first**
   - Only perform follow-up searches if the initial results are insufficient

3. Prefer:
   - Official documentation
   - Reputable news outlets
   - Primary sources over blogs or forums

4. If results conflict:
   - Mention uncertainty
   - Present multiple perspectives when appropriate

## Result Handling

After retrieving results:
- Summarize the findings clearly and accurately
- Attribute facts to their sources when relevant
- Avoid copying large blocks of text verbatim
- If no reliable results are found, say so explicitly

## Follow-ups

If the user’s request is ambiguous:
- Ask a clarifying question **before** searching

If the topic is broad:
- Provide a short summary
- Offer to dig deeper into a specific subtopic

## Safety and Accuracy

- Do not invent facts that are not present in search results
- Do not assume tool errors mean “no information exists”
- Always distinguish between verified information and speculation
