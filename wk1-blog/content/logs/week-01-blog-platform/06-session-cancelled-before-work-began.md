---
title: "Session Cancelled Before Work Began"
weight: 6
agent: "Crush CLI Agent (model: claude-opus-4-5-20251101)"
duration: "6s"
---

## TL;DR - Session insights

- Session was cancelled after just 6 seconds when user aborted the initial file view operation before the agent could begin work
- The prompt contained an elaborate multi-product specification framework with POSA architecture patterns, RFC 2119 keywords, and traceability matrix requirements
- Task scope included creating both technical requirements (3.1_TECH_REQUIREMENTS.md) and system tests (3.2_SYSTEM_TESTS.md) from existing formal requirements
- The detailed prompt structure suggests a methodical software engineering approach with emphasis on industrial-grade specifications

---

## Session Log

* **USER**: ## CONTEXT
  
  We are going to build technical requirements and system tests. You have already discussed with the customer on their needs and captured everything in `design/1.1_NEEDS.md` and the acceptance criteria for these needs in `design/1.2_ACCEPTANCE.md`. You have also spent time translating these two docs into formal requirements captured in `design/2.1_REQUIREMENTS.md` and associated QA specifications in `design/2.2_QA.md`.
  
  ## ROLE
  
  You are an experience architect with many decades of experience in converting formal product requirements into technical requirements for project teams. You pay lot of attention to detail and make sure the level of detail used to specify technical requirements is sufficient for project teams to have all the information to perform code level design and implement it.
  * You a well versed with all the architecture patterns like "Layered architecture", "Pipes and Filters", "Blackboard", "Broker", "Model View Controller (MVC)", "Presentation Abstraction Control (PAC)", "Microkernel" and "Reflection" (all part of the POSA book written by engineers from Siemens AG).
  * You are also well versed with "Service Oriented Architecture (SOA)" and all the standards associated with it that can be leveraged to deploy large scale SOA projects
  * You prefer simplicity over standardization. While you know all possibile architecture and design patterns in the world, you like to keep things as simple as possible (and not any more simpler) because you know that long term stability requires comprehensible architectures that can be easily decomposed into any form or structure and still be usable.
  * While you don't work on code-design you are a big proponent of DRY, YAGNI and SOLID principles.
  
  NOTE: Like all other folks in the team here, you love to use RFC 2119 for requirements keywords and pay attention to the meaning of these keywords.
  
  ## ADDITIONAL INSTRUCTIONS
  
  * Convert each requirement in `design/2.1_REQUIREMENTS.md` into a technical requirement and write it to `design/3.1_TECH_REQUIREMENTS.md`.
    + Technical requirements are architecture and high-level design specification captured in a similar structure as the product requirements you are using as the source.
    + Technical requirements dig deeper into technical decisions and use those decisions to specify technical requirements.
    + Compared to product requirements, technical requirements give full weightage to product requirements, infrastructure requirements, user requirements, experience design requireements, security requirements as well as non-function requirements like performance, resilience, etc. (not all non-functional requirements are listed here. You need to think about all of them and apply)
    + While you are focusing on all aspects of technical requirements (see previous point), you have to make sure they are listed in proper order so that the reader who is responsible for implementing it gets entire context required for implementation by reading adjacent requirements i.e., your technical requirements doc should read like a story, not an overly-formalized and overly-structured specification)
    + Similar to product requirements, you'll have to pick appropriate keywords for requirement IDs. This time, I am not going to prescribe the ID structure. Pick one that is the easiest for a human to read and remember.
    * You must style each requirement as a formally similar to how engineering companies (industrial automation, automotive, building automation, etc.) do it.
    * The file should be structured as follows -
      ```md
      # TECHNICAL REQUIREMENTS
  
      ## PRODUCTS
  
      <INCLUDE THIS SECTION ONLY IF YOU HAVE DECIDED TO BUILD MORE THAN ONE PRODUCT. IF SO, CREATE A NUMBERED LIST OF PRODUCT AND 1-3 LINE DESCRIPTION ABOUT THAT PRODUCT>
  
      ### <FIRST PRODUCT NAME>
  
      <SUMMARY OF THE PRODUCT AND ITS SCOPE AT A TECH REQUIREMENTS LEVEL OF DETAIL>
  
      #### SPECIFICATIONS
  
      * **<3_TO_8_LETTER_PRODUCT_KEYWORD>001** - <TECHNICAL REQUIREMENT>
      * **<3_TO_8_LETTER_PRODUCT_KEYWORD>002** - <TECHNICAL REQUIREMENT>
      * ...
  
      ### <SECOND PRODUCT NAME>
  
      <SUMMARY OF THE PRODUCT AND ITS SCOPE AT A TECH REQUIREMENTS LEVEL OF DETAIL>
  
      #### SPECIFICATIONS
  
      * **<3_TO_8_LETTER_PRODUCT_KEYWORD>001** - <TECHNICAL REQUIREMENT>
      * **<3_TO_8_LETTER_PRODUCT_KEYWORD>002** - <TECHNICAL REQUIREMENT>
      * ...
  
      ## TRACEABILITY MATRIX
  
      <A TABLE MAPPING REQUIREMENT ID TO ONE OR MORE TECHNICAL REQUIREMENTS ID ACROSS PRODUCTS>
  
      ## REFERENCES
  
      <AN OPTIONAL BULLETED LIST OF EXTERNAL DOCUMENTS, WEBSITES, BLOGS, ETC., THAT WAS USED TO BUILD THESE TECHNCIAL REQUIREMENTS>
      ```
  * `design/3.2_SYSTEM_TESTS.md`
    + This file will hold system test specification mapping to each requirement from `design/3.1_TECH_REQUIREMENTS.md`
    + While we are not creating test scripts, the level of details must be deep enough that anyone reading these test specifications should be able to write automation test scripts to test the entire product. The tests must stay true to the actual needs from the user and the formal requriement from the requirements analyst. This information will be fed to an automated system testing LLM agent who will write these tests and run them with almost no supervision. So you are free to make this test specification as long as you want (even 10x-50x larger than the tech requirements spec if required)
    + The file should be structured as follows -
      ```md
      # SYSTEM TESTS
  
      <SHORT SUMMARY OF SCOPE OF SYSTEM TESTS FOR THIS PROJECT>
  
      ## SPECIFICATIONS
  
      * **<3_TO_8_LETTER_PRODUCT_KEYWORD>_TEST_001** - <SYSTEM TEST SPECIFICATION SPECIFIC TO 001 TECH REQUIREMENT OF THE PRODUCT>
        + <OPTIONAL BULLETED LIST OF SUB-SCENARIOS CONNECTED TO THIS TEST SPECIFICATION. EACH SCENARIO WILL BE TRANSLATED INTO A SEPARATE TEST SCRIPT OR TESTING FUNCTION>
      * **<3_TO_8_LETTER_PRODUCT_KEYWORD>_TEST_001** - <SYSTEM TEST SPECIFICATION SPECIFIC TO 001 TECH REQUIREMENT OF THE PRODUCT. LOOKS LIKE THIS SPECIFIC QA SPECIFICATION DOES NOT HAVE SUB-SCENARIOS>
      * ...
      ```
  ## TASK
  
  1. Populate `design/3.1_TECH_REQUIREMENTS.md` document and convert each formal requirement into one or more technical requirements across multiple products (if you decided to build multiple products). Follow the doc template for requirements I've already given you. Technical requirements will be given to a LLM code design agent with minimal human oversight. So you are free to make this spec as large as required.
  2. Use `design/2.1_QA.md` to build `design/3.2_SYSTEM_TESTS.md`. You may refer `design/1.1_NEEDS.md` to undersstand the context and `design/1.2_ACCEPTANCE.md` for acceptance tests that customer will run to confirm that the needs are met. Each system test spec must map to the exact requirement from `design/3.1_TECH_REQUIREMENTS.md`. So, if you have specified requirements for multiple products, you will build system test specs for multiple products too.
  
  

* **TOOL-CALL**: view `{}`

* **CLI-MESSAGE**: `Tool execution canceled by user`
