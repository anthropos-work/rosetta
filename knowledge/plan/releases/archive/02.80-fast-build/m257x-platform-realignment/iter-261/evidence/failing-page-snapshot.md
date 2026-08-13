# Instructions

- Following Playwright test failed.
- Explain why, be concise, respect Playwright best practices.
- Provide a snippet of code with the fix, if possible.

# Test info

- Name: workforce-succession.spec.ts >> @pt:pt-workforce-succession — manager succession / at-risk >> a manager sees the succession candidates and at-risk signals render
- Location: tests/workforce-succession.spec.ts:35:7

# Error details

```
Error: the succession/at-risk projection names the org's seeded hero (Pat Ellis) — which is what makes it THIS org's projection rather than any populated org's

the succession/at-risk projection names the org's seeded hero (Pat Ellis) — which is what makes it THIS org's projection rather than any populated org's

expect(received).toBeGreaterThan(expected)

Expected: > 0
Received:   0

Call Log:
- Timeout 15000ms exceeded while waiting on the predicate
```

# Page snapshot

```yaml
- generic [active] [ref=e1]:
  - alert [ref=e2]
  - generic [ref=e4]:
    - banner [ref=e6]:
      - generic [ref=e7]:
        - link "Anthropos Workforce" [ref=e9] [cursor=pointer]:
          - /url: /home
          - generic [ref=e10]:
            - img "Anthropos" [ref=e11]
            - generic [ref=e13]: Workforce
        - button "Search... ✦ AI Search" [ref=e14] [cursor=pointer]:
          - img [ref=e15]
          - generic [ref=e17]: Search...
          - generic [ref=e18]: ✦ AI Search
        - generic [ref=e19]:
          - button "Meridian Labs · Workforce" [ref=e21] [cursor=pointer]:
            - img "Meridian Labs" [ref=e23]
            - generic [ref=e24]:
              - generic [ref=e25]: Meridian Labs
              - generic [ref=e26]: Workforce
            - img [ref=e27]
          - separator [ref=e29]
          - button [ref=e31] [cursor=pointer]:
            - img [ref=e33]
          - separator [ref=e41]
          - button "MR Morgan" [ref=e43] [cursor=pointer]:
            - img "MR" [ref=e45]
            - generic [ref=e46]: Morgan
            - img [ref=e47]
    - generic [ref=e49]:
      - generic [ref=e50]:
        - complementary [ref=e51]:
          - generic [ref=e53]:
            - generic [ref=e54]:
              - menu [ref=e55]:
                - menuitem "Home" [ref=e56] [cursor=pointer]:
                  - link [ref=e58]:
                    - /url: /home
                    - img [ref=e59]
                  - link "Home" [ref=e63]:
                    - /url: /home
                - menuitem "Career Profile" [ref=e64] [cursor=pointer]:
                  - link [ref=e66]:
                    - /url: /profile
                    - img [ref=e67]
                  - link "Career Profile" [ref=e71]:
                    - /url: /profile
                - menuitem "My Skills" [ref=e72] [cursor=pointer]:
                  - link [ref=e74]:
                    - /url: /profile/skills
                    - img [ref=e75]
                  - link "My Skills" [ref=e79]:
                    - /url: /profile/skills
                - menuitem "My Activities" [ref=e80] [cursor=pointer]:
                  - link [ref=e82]:
                    - /url: /profile/activities
                    - img [ref=e83]
                  - link "My Activities" [ref=e87]:
                    - /url: /profile/activities
              - generic [ref=e88]:
                - generic [ref=e90]: Content Library
                - menu [ref=e91]:
                  - menuitem "AI Simulations" [ref=e92] [cursor=pointer]:
                    - link [ref=e94]:
                      - /url: /library/ai-simulations
                      - img [ref=e95]
                    - link "AI Simulations" [ref=e99]:
                      - /url: /library/ai-simulations
                  - menuitem "Skill Paths" [ref=e100] [cursor=pointer]:
                    - link [ref=e102]:
                      - /url: /library/skill-paths
                      - img [ref=e103]
                    - link "Skill Paths" [ref=e107]:
                      - /url: /library/skill-paths
                  - menuitem "AI Academy" [ref=e108] [cursor=pointer]:
                    - link [ref=e110]:
                      - /url: http://localhost:23077
                      - img [ref=e111]
                    - link "AI Academy" [ref=e115]:
                      - /url: http://localhost:23077
              - generic [ref=e116]:
                - generic [ref=e117]:
                  - generic [ref=e118]: Organization
                  - button "Expand all" [ref=e119] [cursor=pointer]
                - button "Map" [ref=e121] [cursor=pointer]:
                  - img [ref=e122]
                  - generic [ref=e124]: Map
                  - img [ref=e125]
                - button "Customize" [ref=e128] [cursor=pointer]:
                  - img [ref=e129]
                  - generic [ref=e131]: Customize
                  - img [ref=e132]
                - button "Assign" [ref=e135] [cursor=pointer]:
                  - img [ref=e136]
                  - generic [ref=e138]: Assign
                  - img [ref=e139]
                - button "Track & Verify" [ref=e142] [cursor=pointer]:
                  - img [ref=e143]
                  - generic [ref=e145]: Track & Verify
                  - img [ref=e146]
                - button "Intelligence" [ref=e149] [cursor=pointer]:
                  - img [ref=e150]
                  - generic [ref=e152]: Intelligence
                  - img [ref=e153]
                - separator [ref=e155]
                - menu [ref=e156]:
                  - menuitem "Talk to Data" [ref=e157] [cursor=pointer]:
                    - link [ref=e159]:
                      - /url: /enterprise/talk-to-data
                      - img [ref=e160]
                    - link "Talk to Data" [ref=e164]:
                      - /url: /enterprise/talk-to-data
                - generic [ref=e166]:
                  - link "Enterprise Settings" [ref=e167] [cursor=pointer]:
                    - /url: /enterprise/settings
                    - img [ref=e168]
                    - generic [ref=e170]: Enterprise Settings
                  - button "Enterprise Settings" [ref=e171] [cursor=pointer]:
                    - img [ref=e172]
            - menu [ref=e176]:
              - menuitem "Help" [ref=e177] [cursor=pointer]:
                - img [ref=e180]
                - generic [ref=e184]: Help
        - button [ref=e185] [cursor=pointer]:
          - img [ref=e186]
      - main [ref=e189]:
        - main [ref=e191]:
          - heading "Succession Planning" [level=1] [ref=e192]
          - generic [ref=e193]:
            - alert [ref=e194]:
              - img "check-circle" [ref=e196]:
                - img [ref=e197]
              - generic [ref=e199]:
                - generic [ref=e200]: Good data coverage
                - generic [ref=e201]: Re-run AI Interviews every 6 months to keep flight risk signals up to date.
            - generic [ref=e202]:
              - generic [ref=e205]:
                - generic [ref=e206]:
                  - generic [ref=e208]: Critical roles
                  - generic [ref=e210]: "6"
                - text: risk ≥ 50
              - generic [ref=e213]:
                - generic [ref=e214]:
                  - generic [ref=e216]: At-risk people
                  - generic [ref=e218]: "4"
                - text: score ≥ 40
              - generic [ref=e222]:
                - generic [ref=e224]: Top talents (ready)
                - generic [ref=e226]: "0"
              - generic [ref=e230]:
                - generic [ref=e231]: Data confidence
                - generic [ref=e232]: full
                - generic [ref=e233]: Skills 100% · Sims 100% · Interview 29%
            - generic [ref=e234]:
              - generic [ref=e237]: Roles by risk
              - generic [ref=e239]:
                - generic [ref=e243] [cursor=pointer]:
                  - generic [ref=e244]:
                    - generic [ref=e245]:
                      - heading "Advanced Analytics Specialist" [level=5] [ref=e246]
                      - generic [ref=e247]: 1 incumbent · 10 required skills
                    - generic [ref=e248]: MEDIUM
                  - generic [ref=e249]:
                    - progressbar [ref=e250]
                    - generic [ref=e254]: risk 68
                  - generic [ref=e255]:
                    - generic [ref=e256]: 0 ready
                    - generic [ref=e257]: 0 dev
                    - generic [ref=e258]: Avg fit 5%
                  - generic [ref=e259]:
                    - generic [ref=e260]: Critical Thinking Fundamentals
                    - generic [ref=e261]: Data Analysis
                    - generic [ref=e262]: Data Cleaning and Preprocessing
                    - generic [ref=e263]: Data Governance and Quality
                    - generic [ref=e264]: Data Mining
                    - generic [ref=e265]: +5 more
                - generic [ref=e269] [cursor=pointer]:
                  - generic [ref=e270]:
                    - generic [ref=e271]:
                      - heading "Engineering Manager" [level=5] [ref=e272]
                      - generic [ref=e273]: 1 incumbent · 10 required skills
                    - generic [ref=e274]: MEDIUM
                  - generic [ref=e275]:
                    - progressbar [ref=e276]
                    - generic [ref=e280]: risk 68
                  - generic [ref=e281]:
                    - generic [ref=e282]: 0 ready
                    - generic [ref=e283]: 0 dev
                    - generic [ref=e284]: Avg fit 44%
                  - generic [ref=e285]:
                    - generic [ref=e286]: Budgeting
                    - generic [ref=e287]: Cross-Functional Collaboration
                    - generic [ref=e288]: Leadership and Teamwork
                    - generic [ref=e289]: Mentorship and Coaching
                    - generic [ref=e290]: Performance Management Fundamentals
                    - generic [ref=e291]: +5 more
                - generic [ref=e295] [cursor=pointer]:
                  - generic [ref=e296]:
                    - generic [ref=e297]:
                      - heading "Administrative Coordinator" [level=5] [ref=e298]
                      - generic [ref=e299]: 1 incumbent · 10 required skills
                    - generic [ref=e300]: MEDIUM
                  - generic [ref=e301]:
                    - progressbar [ref=e302]
                    - generic [ref=e306]: risk 68
                  - generic [ref=e307]:
                    - generic [ref=e308]: 0 ready
                    - generic [ref=e309]: 0 dev
                    - generic [ref=e310]: Avg fit 0%
                  - generic [ref=e311]:
                    - generic [ref=e312]: Adapting to Changing Work Environments
                    - generic [ref=e313]: Administrative Role Preparation
                    - generic [ref=e314]: Attention to Detail
                    - generic [ref=e315]: Communication
                    - generic [ref=e316]: Data Organization
                    - generic [ref=e317]: +5 more
                - generic [ref=e321] [cursor=pointer]:
                  - generic [ref=e322]:
                    - generic [ref=e323]:
                      - heading "Admin & HR Coordinator" [level=5] [ref=e324]
                      - generic [ref=e325]: 1 incumbent · 10 required skills
                    - generic [ref=e326]: MEDIUM
                  - generic [ref=e327]:
                    - progressbar [ref=e328]
                    - generic [ref=e332]: risk 68
                  - generic [ref=e333]:
                    - generic [ref=e334]: 0 ready
                    - generic [ref=e335]: 0 dev
                    - generic [ref=e336]: Avg fit 3%
                  - generic [ref=e337]:
                    - generic [ref=e338]: Communication
                    - generic [ref=e339]: Compliance Management
                    - generic [ref=e340]: Document Management Software Usage
                    - generic [ref=e341]: Employee Collaboration
                    - generic [ref=e342]: Employee Relations Fundamentals
                    - generic [ref=e343]: +5 more
                - generic [ref=e347] [cursor=pointer]:
                  - generic [ref=e348]:
                    - generic [ref=e349]:
                      - heading "Adjunct Professor - Strategy, Leadership And People Department" [level=5] [ref=e350]
                      - generic [ref=e351]: 2 incumbents · 10 required skills
                    - generic [ref=e352]: MEDIUM
                  - generic [ref=e353]:
                    - progressbar [ref=e354]
                    - generic [ref=e358]: risk 54
                  - generic [ref=e359]:
                    - generic [ref=e360]: 0 ready
                    - generic [ref=e361]: 0 dev
                    - generic [ref=e362]: Avg fit 0%
                  - generic [ref=e363]:
                    - generic [ref=e364]: Analysis Skills
                    - generic [ref=e365]: Curriculum Design and Development
                    - generic [ref=e366]: Leadership and Change Management
                    - generic [ref=e367]: Leadership Communication
                    - generic [ref=e368]: Organizational Culture Fundamentals
                    - generic [ref=e369]: +5 more
                - generic [ref=e373] [cursor=pointer]:
                  - generic [ref=e374]:
                    - generic [ref=e375]:
                      - heading "Advanced Analytics and Business Intelligence Manager" [level=5] [ref=e376]
                      - generic [ref=e377]: 2 incumbents · 10 required skills
                    - generic [ref=e378]: MEDIUM
                  - generic [ref=e379]:
                    - progressbar [ref=e380]
                    - generic [ref=e384]: risk 54
                  - generic [ref=e385]:
                    - generic [ref=e386]: 0 ready
                    - generic [ref=e387]: 0 dev
                    - generic [ref=e388]: Avg fit 8%
                  - generic [ref=e389]:
                    - generic [ref=e390]: Business Intelligence
                    - generic [ref=e391]: Communication
                    - generic [ref=e392]: Cross-Functional Team Leadership
                    - generic [ref=e393]: Data Analysis
                    - generic [ref=e394]: Data Visualization
                    - generic [ref=e395]: +5 more
                - generic [ref=e399] [cursor=pointer]:
                  - generic [ref=e400]:
                    - generic [ref=e401]:
                      - heading "Administrative Assistant" [level=5] [ref=e402]
                      - generic [ref=e403]: 3 incumbents · 10 required skills
                    - generic [ref=e404]: LOW
                  - generic [ref=e405]:
                    - progressbar [ref=e406]
                    - generic [ref=e410]: risk 45
                  - generic [ref=e411]:
                    - generic [ref=e412]: 0 ready
                    - generic [ref=e413]: 0 dev
                    - generic [ref=e414]: Avg fit 6%
                  - generic [ref=e415]:
                    - generic [ref=e416]: Attention to Detail
                    - generic [ref=e417]: Customer Service Skills
                    - generic [ref=e418]: Interpersonal Skills
                    - generic [ref=e419]: Job Description Analysis
                    - generic [ref=e420]: Microsoft Office Suite
                    - generic [ref=e421]: +5 more
                - generic [ref=e425] [cursor=pointer]:
                  - generic [ref=e426]:
                    - generic [ref=e427]:
                      - heading "Administrator" [level=5] [ref=e428]
                      - generic [ref=e429]: 3 incumbents · 10 required skills
                    - generic [ref=e430]: LOW
                  - generic [ref=e431]:
                    - progressbar [ref=e432]
                    - generic [ref=e436]: risk 45
                  - generic [ref=e437]:
                    - generic [ref=e438]: 0 ready
                    - generic [ref=e439]: 0 dev
                    - generic [ref=e440]: Avg fit 4%
                  - generic [ref=e441]:
                    - generic [ref=e442]: Attention to Detail
                    - generic [ref=e443]: Budgeting
                    - generic [ref=e444]: Documentation and Record Keeping
                    - generic [ref=e445]: Google Docs
                    - generic [ref=e446]: Microsoft Office Suite
                    - generic [ref=e447]: +5 more
                - generic [ref=e451] [cursor=pointer]:
                  - generic [ref=e452]:
                    - generic [ref=e453]:
                      - heading "DevOps Engineer" [level=5] [ref=e454]
                      - generic [ref=e455]: 3 incumbents · 10 required skills
                    - generic [ref=e456]: LOW
                  - generic [ref=e457]:
                    - progressbar [ref=e458]
                    - generic [ref=e462]: risk 45
                  - generic [ref=e463]:
                    - generic [ref=e464]: 0 ready
                    - generic [ref=e465]: 0 dev
                    - generic [ref=e466]: Avg fit 51%
                  - generic [ref=e467]:
                    - generic [ref=e468]: Agile Methodologies
                    - generic [ref=e469]: Automation Tools
                    - generic [ref=e470]: Containerization (Docker, Kubernetes)
                    - generic [ref=e471]: Continuous Integration/Continuous Deployment (CI/CD) Pipelines
                    - generic [ref=e472]: Infrastructure-as-Code (IaC)
                    - generic [ref=e473]: +5 more
                - generic [ref=e477] [cursor=pointer]:
                  - generic [ref=e478]:
                    - generic [ref=e479]:
                      - heading "Administrative Virtual Assistant" [level=5] [ref=e480]
                      - generic [ref=e481]: 4 incumbents · 10 required skills
                    - generic [ref=e482]: LOW
                  - generic [ref=e483]:
                    - progressbar [ref=e484]
                    - generic [ref=e488]: risk 45
                  - generic [ref=e489]:
                    - generic [ref=e490]: 0 ready
                    - generic [ref=e491]: 0 dev
                    - generic [ref=e492]: Avg fit 2%
                  - generic [ref=e493]:
                    - generic [ref=e494]: Communication Abilities
                    - generic [ref=e495]: Confidentiality and Privacy Protection
                    - generic [ref=e496]: Customer Service Skills
                    - generic [ref=e497]: Digital and Virtual Collaboration Tools
                    - generic [ref=e498]: Email System Support
                    - generic [ref=e499]: +5 more
                - generic [ref=e503] [cursor=pointer]:
                  - generic [ref=e504]:
                    - generic [ref=e505]:
                      - heading "Ad Tech Engineer" [level=5] [ref=e506]
                      - generic [ref=e507]: 4 incumbents · 10 required skills
                    - generic [ref=e508]: LOW
                  - generic [ref=e509]:
                    - progressbar [ref=e510]
                    - generic [ref=e514]: risk 45
                  - generic [ref=e515]:
                    - generic [ref=e516]: 0 ready
                    - generic [ref=e517]: 0 dev
                    - generic [ref=e518]: Avg fit 0%
                  - generic [ref=e519]:
                    - generic [ref=e520]: Ad Management Platforms
                    - generic [ref=e521]: API Integration
                    - generic [ref=e522]: Cross-Functional Collaboration
                    - generic [ref=e523]: Data Privacy Compliance
                    - generic [ref=e524]: JavaScript
                    - generic [ref=e525]: +5 more
                - generic [ref=e529] [cursor=pointer]:
                  - generic [ref=e530]:
                    - generic [ref=e531]:
                      - heading "Administration officer" [level=5] [ref=e532]
                      - generic [ref=e533]: 3 incumbents · 10 required skills
                    - generic [ref=e534]: LOW
                  - generic [ref=e535]:
                    - progressbar [ref=e536]
                    - generic [ref=e540]: risk 45
                  - generic [ref=e541]:
                    - generic [ref=e542]: 0 ready
                    - generic [ref=e543]: 0 dev
                    - generic [ref=e544]: Avg fit 4%
                  - generic [ref=e545]:
                    - generic [ref=e546]: Customer Service Skills
                    - generic [ref=e547]: Data Organization and Recording
                    - generic [ref=e548]: Documentation Development and Management
                    - generic [ref=e549]: Email Management
                    - generic [ref=e550]: Microsoft Excel
                    - generic [ref=e551]: +5 more
            - generic [ref=e552]:
              - generic [ref=e555]: Readiness Explorer
              - generic [ref=e557]:
                - generic [ref=e558]:
                  - generic [ref=e559]:
                    - generic "Advanced Analytics Specialist (1 in role · 0 candidates)" [ref=e560]:
                      - text: Advanced Analytics Specialist (1 in role · 0 candidates)
                      - combobox [ref=e561]
                    - img "down" [ref=e563]:
                      - img [ref=e564]
                  - generic [ref=e566]:
                    - searchbox "Filter candidates…" [ref=e568]
                    - button "search" [ref=e570] [cursor=pointer]:
                      - img "search" [ref=e572]:
                        - img [ref=e573]
                - generic [ref=e575]:
                  - generic [ref=e576]: "Required:"
                  - generic [ref=e577]:
                    - text: Critical Thinking Fundamentals
                    - generic [ref=e578]: L4
                  - generic [ref=e579]:
                    - text: Data Analysis
                    - generic [ref=e580]: L5
                  - generic [ref=e581]:
                    - text: Data Cleaning and Preprocessing
                    - generic [ref=e582]: L5
                  - generic [ref=e583]:
                    - text: Data Governance and Quality
                    - generic [ref=e584]: L4
                  - generic [ref=e585]:
                    - text: Data Mining
                    - generic [ref=e586]: L4
                  - generic [ref=e587]:
                    - text: Data Visualization
                    - generic [ref=e588]: L5
                  - generic [ref=e589]:
                    - text: Machine Learning Application
                    - generic [ref=e590]: L5
                  - generic [ref=e591]:
                    - text: Python
                    - generic [ref=e592]: L5
                  - generic [ref=e593]:
                    - text: SQL
                    - generic [ref=e594]: L5
                  - generic [ref=e595]:
                    - text: Statistical and Predictive Modeling
                    - generic [ref=e596]: L5
                - generic [ref=e597]:
                  - img "No data" [ref=e599]
                  - generic [ref=e610]: No candidates match
            - generic [ref=e611]:
              - generic [ref=e614]: People at risk
              - generic [ref=e618]:
                - table [ref=e622]:
                  - rowgroup [ref=e628]:
                    - row "Name Score Signals Interview" [ref=e629]:
                      - columnheader "Name" [ref=e630]
                      - columnheader "Score" [ref=e631] [cursor=pointer]:
                        - generic [ref=e632]:
                          - generic [ref=e633]: Score
                          - generic [ref=e635]:
                            - img [ref=e636]:
                              - img [ref=e637]
                            - img [ref=e639]:
                              - img [ref=e640]
                      - columnheader "Signals" [ref=e642]
                      - columnheader "Interview" [ref=e643]
                  - rowgroup [ref=e644]:
                    - 'row "Hassan Larsen Administrator 95 Low wellbeing (D) Negative sentiment from AI Interview 4 attention signals from AI Interview Exit intent: non soddisfatto Skills not aligned with role (fit 11%) yes" [ref=e645]':
                      - cell "Hassan Larsen Administrator" [ref=e646]:
                        - generic [ref=e647]:
                          - strong [ref=e649]: Hassan Larsen
                          - generic [ref=e650]: Administrator
                      - cell "95" [ref=e651]
                      - 'cell "Low wellbeing (D) Negative sentiment from AI Interview 4 attention signals from AI Interview Exit intent: non soddisfatto Skills not aligned with role (fit 11%)" [ref=e652]':
                        - generic [ref=e653]:
                          - generic [ref=e654]: Low wellbeing (D)
                          - generic [ref=e655]: Negative sentiment from AI Interview
                          - generic [ref=e656]: 4 attention signals from AI Interview
                          - generic [ref=e657]: "Exit intent: non soddisfatto"
                          - generic [ref=e658]: Skills not aligned with role (fit 11%)
                      - cell "yes" [ref=e659]:
                        - generic [ref=e660]: "yes"
                    - 'row "Chloe Mensah Advanced Analytics Specialist 50 Rare skill held only by this person In fragile role: Advanced Analytics Specialist Skills not aligned with role (fit 5%) no" [ref=e661]':
                      - cell "Chloe Mensah Advanced Analytics Specialist" [ref=e662]:
                        - generic [ref=e663]:
                          - strong [ref=e665]: Chloe Mensah
                          - generic [ref=e666]: Advanced Analytics Specialist
                      - cell "50" [ref=e667]
                      - 'cell "Rare skill held only by this person In fragile role: Advanced Analytics Specialist Skills not aligned with role (fit 5%)" [ref=e668]':
                        - generic [ref=e669]:
                          - generic [ref=e670]: Rare skill held only by this person
                          - generic [ref=e671]: "In fragile role: Advanced Analytics Specialist"
                          - generic [ref=e672]: Skills not aligned with role (fit 5%)
                      - cell "no" [ref=e673]:
                        - generic [ref=e674]: "no"
                    - 'row "Morgan Reyes Engineering Manager 50 Rare skill held only by this person In fragile role: Engineering Manager Skills not aligned with role (fit 44%) yes" [ref=e675]':
                      - cell "Morgan Reyes Engineering Manager" [ref=e676]:
                        - generic [ref=e677]:
                          - strong [ref=e679]: Morgan Reyes
                          - generic [ref=e680]: Engineering Manager
                      - cell "50" [ref=e681]
                      - 'cell "Rare skill held only by this person In fragile role: Engineering Manager Skills not aligned with role (fit 44%)" [ref=e682]':
                        - generic [ref=e683]:
                          - generic [ref=e684]: Rare skill held only by this person
                          - generic [ref=e685]: "In fragile role: Engineering Manager"
                          - generic [ref=e686]: Skills not aligned with role (fit 44%)
                      - cell "yes" [ref=e687]:
                        - generic [ref=e688]: "yes"
                    - 'row "Diego Lindqvist Advanced Analytics and Business Intelligence Manager 50 Rare skill held only by this person In fragile role: Advanced Analytics and Business Intelligence Manager Skills not aligned with role (fit 15%) no" [ref=e689]':
                      - cell "Diego Lindqvist Advanced Analytics and Business Intelligence Manager" [ref=e690]:
                        - generic [ref=e691]:
                          - strong [ref=e693]: Diego Lindqvist
                          - generic [ref=e694]: Advanced Analytics and Business Intelligence Manager
                      - cell "50" [ref=e695]
                      - 'cell "Rare skill held only by this person In fragile role: Advanced Analytics and Business Intelligence Manager Skills not aligned with role (fit 15%)" [ref=e696]':
                        - generic [ref=e697]:
                          - generic [ref=e698]: Rare skill held only by this person
                          - generic [ref=e699]: "In fragile role: Advanced Analytics and Business Intelligence Manager"
                          - generic [ref=e700]: Skills not aligned with role (fit 15%)
                      - cell "no" [ref=e701]:
                        - generic [ref=e702]: "no"
                    - row "Ava Romano Administrative Virtual Assistant 35 Rare skill held only by this person Skills not aligned with role (fit 8%) no" [ref=e703]:
                      - cell "Ava Romano Administrative Virtual Assistant" [ref=e704]:
                        - generic [ref=e705]:
                          - strong [ref=e707]: Ava Romano
                          - generic [ref=e708]: Administrative Virtual Assistant
                      - cell "35" [ref=e709]
                      - cell "Rare skill held only by this person Skills not aligned with role (fit 8%)" [ref=e710]:
                        - generic [ref=e711]:
                          - generic [ref=e712]: Rare skill held only by this person
                          - generic [ref=e713]: Skills not aligned with role (fit 8%)
                      - cell "no" [ref=e714]:
                        - generic [ref=e715]: "no"
                    - row "Layla Kovac Administration officer 35 Rare skill held only by this person Skills not aligned with role (fit 3%) no" [ref=e716]:
                      - cell "Layla Kovac Administration officer" [ref=e717]:
                        - generic [ref=e718]:
                          - strong [ref=e720]: Layla Kovac
                          - generic [ref=e721]: Administration officer
                      - cell "35" [ref=e722]
                      - cell "Rare skill held only by this person Skills not aligned with role (fit 3%)" [ref=e723]:
                        - generic [ref=e724]:
                          - generic [ref=e725]: Rare skill held only by this person
                          - generic [ref=e726]: Skills not aligned with role (fit 3%)
                      - cell "no" [ref=e727]:
                        - generic [ref=e728]: "no"
                    - row "Arjun Larsen Administrative Assistant 35 Rare skill held only by this person Skills not aligned with role (fit 10%) no" [ref=e729]:
                      - cell "Arjun Larsen Administrative Assistant" [ref=e730]:
                        - generic [ref=e731]:
                          - strong [ref=e733]: Arjun Larsen
                          - generic [ref=e734]: Administrative Assistant
                      - cell "35" [ref=e735]
                      - cell "Rare skill held only by this person Skills not aligned with role (fit 10%)" [ref=e736]:
                        - generic [ref=e737]:
                          - generic [ref=e738]: Rare skill held only by this person
                          - generic [ref=e739]: Skills not aligned with role (fit 10%)
                      - cell "no" [ref=e740]:
                        - generic [ref=e741]: "no"
                    - 'row "Chloe Nakamura Administrative Coordinator 25 In fragile role: Administrative Coordinator Skills not aligned with role (fit 0%) yes" [ref=e742]':
                      - cell "Chloe Nakamura Administrative Coordinator" [ref=e743]:
                        - generic [ref=e744]:
                          - strong [ref=e746]: Chloe Nakamura
                          - generic [ref=e747]: Administrative Coordinator
                      - cell "25" [ref=e748]
                      - 'cell "In fragile role: Administrative Coordinator Skills not aligned with role (fit 0%)" [ref=e749]':
                        - generic [ref=e750]:
                          - generic [ref=e751]: "In fragile role: Administrative Coordinator"
                          - generic [ref=e752]: Skills not aligned with role (fit 0%)
                      - cell "yes" [ref=e753]:
                        - generic [ref=e754]: "yes"
                    - 'row "Omar Moreau Adjunct Professor - Strategy, Leadership And People Department 25 In fragile role: Adjunct Professor - Strategy, Leadership And People Department Skills not aligned with role (fit 0%) no" [ref=e755]':
                      - cell "Omar Moreau Adjunct Professor - Strategy, Leadership And People Department" [ref=e756]:
                        - generic [ref=e757]:
                          - strong [ref=e759]: Omar Moreau
                          - generic [ref=e760]: Adjunct Professor - Strategy, Leadership And People Department
                      - cell "25" [ref=e761]
                      - 'cell "In fragile role: Adjunct Professor - Strategy, Leadership And People Department Skills not aligned with role (fit 0%)" [ref=e762]':
                        - generic [ref=e763]:
                          - generic [ref=e764]: "In fragile role: Adjunct Professor - Strategy, Leadership And People Department"
                          - generic [ref=e765]: Skills not aligned with role (fit 0%)
                      - cell "no" [ref=e766]:
                        - generic [ref=e767]: "no"
                    - 'row "Ethan Silva Adjunct Professor - Strategy, Leadership And People Department 25 In fragile role: Adjunct Professor - Strategy, Leadership And People Department Skills not aligned with role (fit 0%) no" [ref=e768]':
                      - cell "Ethan Silva Adjunct Professor - Strategy, Leadership And People Department" [ref=e769]:
                        - generic [ref=e770]:
                          - strong [ref=e772]: Ethan Silva
                          - generic [ref=e773]: Adjunct Professor - Strategy, Leadership And People Department
                      - cell "25" [ref=e774]
                      - 'cell "In fragile role: Adjunct Professor - Strategy, Leadership And People Department Skills not aligned with role (fit 0%)" [ref=e775]':
                        - generic [ref=e776]:
                          - generic [ref=e777]: "In fragile role: Adjunct Professor - Strategy, Leadership And People Department"
                          - generic [ref=e778]: Skills not aligned with role (fit 0%)
                      - cell "no" [ref=e779]:
                        - generic [ref=e780]: "no"
                - list [ref=e781]:
                  - listitem "Previous Page" [ref=e782]:
                    - button "left" [disabled] [ref=e783]:
                      - img "left" [ref=e784]:
                        - img [ref=e785]
                  - listitem "1" [ref=e787] [cursor=pointer]:
                    - generic [ref=e788]: "1"
                  - listitem "2" [ref=e789] [cursor=pointer]:
                    - generic [ref=e790]: "2"
                  - listitem "3" [ref=e791] [cursor=pointer]:
                    - generic [ref=e792]: "3"
                  - listitem "Next Page" [ref=e793] [cursor=pointer]:
                    - button "right" [ref=e794]:
                      - img "right" [ref=e795]:
                        - img [ref=e796]
            - generic [ref=e798]:
              - generic [ref=e801]: Top talents ready for promotion
              - generic [ref=e803]:
                - img "No data" [ref=e805]
                - generic [ref=e816]: No ready successors identified
```

# Test source

```ts
  1   | /**
  2   |  * pt-workforce-succession — the manager succession / at-risk Playthrough
  3   |  * (manifest use case workforce-intelligence.talent-pool.UC1; M204 manager vantage).
  4   |  *
  5   |  * @pt:pt-workforce-succession
  6   |  * @pt-mutation: READ-ONLY
  7   |  * @pt-mutation-evidence: reads the succession surface; navigation + assertions only
  8   |  *
  9   |  * WHAT IT PROVES: a manager (Morgan / pt-manager) opens the succession view and the projection
  10  |  * ran over **her own org's** people — the seeded hero's job role is one of the computed key-role
  11  |  * cards, and the hero herself appears in the succession/at-risk table with that role — not an
  12  |  * empty/placeholder surface and not another tenant's projection.
  13  |  *
  14  |  *   login as the manager  →  /enterprise/workforce/succession  →  assert the succession + at-risk
  15  |  *   sections rendered  →  assert the seeded role has a key-role card + the seeded hero has a row.
  16  |  *
  17  |  * SHARPENED at v2.8 M256 iter-14. The old final (/ready/i, /at.?risk/i, rows > 0) is STRUCTURAL
  18  |  * and measurably true of a second seeded org, so it proved a projection computed, not whose.
  19  |  *
  20  |  * READ/monitoring flow (no mutation). Succession/at-risk are COMPUTED PROJECTIONS
  21  |  * (trajectory-aware), so per §5.2/P2 + the M204 spec-notes we assert their PRESENCE / STRUCTURE
  22  |  * (the surface computed + rendered real signals), never exact successor identities / risk
  23  |  * values that vary with the seed.
  24  |  */
  25  | import { test, expect } from '@playwright/test';
  26  | import { resolveStackEnv } from '../lib/stack-env';
  27  | import { loginAsHero } from '../lib/hero-login';
  28  | import { SuccessionPage } from '../lib/succession-page';
  29  | import { SUCCESSION_URL } from '../lib/url-shapes';
  30  | import { PT_EMPLOYEE } from '../lib/seed-facts';
  31  | 
  32  | const HERO_SEAT = 'pt-manager';
  33  | 
  34  | test.describe('@pt:pt-workforce-succession — manager succession / at-risk', () => {
  35  |   test('a manager sees the succession candidates and at-risk signals render', async ({ page }) => {
  36  |     const env = resolveStackEnv();
  37  | 
  38  |     const landedUrl = await loginAsHero(page, {
  39  |       appBaseUrl: env.appBaseUrl,
  40  |       fapiBaseUrl: env.fapiBaseUrl,
  41  |       identityKey: HERO_SEAT,
  42  |       landingPath: '/enterprise/workforce',
  43  |       // ANTI-NETWORKIDLE (v2.8 M256 iter-03): the cockpit-login default is 'networkidle', which next-web's
  44  |       // long-poll connections never reach for the right reason — measured 2854 ms vs 423 ms for the SAME
  45  |       // navigation on this surface. domcontentloaded settles the nav; the assertions below auto-retry.
  46  |       waitUntil: 'domcontentloaded',
  47  |     });
  48  |     expect(landedUrl, 'the manager is logged in (not a /login bounce)').not.toMatch(/\/login\b/);
  49  | 
  50  |     const sp = new SuccessionPage(page, env.appBaseUrl);
  51  |     await sp.goto();
  52  | 
  53  |     // intermediate (label: succession-reachable) — the succession view landed.
  54  |     expect(sp.currentUrl(), 'the manager is on the succession view').toMatch(SUCCESSION_URL);
  55  |     await expect(sp.pageHeading(), 'the Succession Planning view rendered').toBeVisible();
  56  | 
  57  |     // intermediate (label: projection-renders) — the succession signal, the at-risk section, and real
  58  |     // rows. THIS USED TO BE THE FINAL, demoted at v2.8 M256 iter-14: /ready/i, /at.?risk/i and "the
  59  |     // table has rows" are STRUCTURAL — they hold for any org whose projection computed at all, so the
  60  |     // Playthrough would have passed on another tenant's succession view. Kept as the intermediate it is.
  61  |     await expect(
  62  |       sp.successionSignal(),
  63  |       'the succession-candidate structure renders (ready talents / successor signal)',
  64  |     ).toBeVisible();
  65  |     await expect(
  66  |       sp.atRiskSignal(),
  67  |       'the at-risk section renders (the trajectory-aware at-risk projection)',
  68  |     ).toBeVisible();
  69  |     await expect(
  70  |       sp.roleCandidateRows().first(),
  71  |       'the role→succession-candidate table rendered real rows (not an empty/placeholder)',
  72  |     ).toBeVisible();
  73  |     const rows = await sp.roleCandidateRows().count();
  74  |     expect(rows, 'the succession projection tabulated real roles for the org').toBeGreaterThan(0);
  75  | 
  76  |     // final, part 1 — the projection ran over THIS org's key roles: the seeded hero's own job role is one
  77  |     // of the role cards the view computes. Measured (iter-14 Phase A): `DevOps Engineer` is a heading here
  78  |     // and the contrast org's heading set is entirely different (Data Analyst, AI Designer, …), so this is
  79  |     // false for another tenant while "a role card exists" is true for all of them.
  80  |     await expect(
  81  |       sp.keyRoleCard(PT_EMPLOYEE.role),
  82  |       `the projection computed a card for the org's own seeded role ("${PT_EMPLOYEE.role}")`,
  83  |     ).toBeVisible();
  84  | 
  85  |     // final, part 2 — and it reached the org's own seeded PERSON: her row carries her seeded role, so the
  86  |     // projection is about her and not about a generated look-alike. Row-scoped (`has:`, not `hasText:` —
  87  |     // iter-13 D63). Measured: `"Pat Ellis / DevOps Engineer / 40 / Rare skill held only by this person /
  88  |     // In fragile role: DevOps Engineer"`; **0** rows for the contrast manager.
  89  |     const heroRow = sp.talentRow(PT_EMPLOYEE.name);
  90  |     await expect
  91  |       .poll(() => heroRow.count(), {
  92  |         message:
  93  |           `the succession/at-risk projection names the org's seeded hero (${PT_EMPLOYEE.name}) — which is ` +
  94  |           `what makes it THIS org's projection rather than any populated org's`,
  95  |       })
> 96  |       .toBeGreaterThan(0);
      |        ^ Error: the succession/at-risk projection names the org's seeded hero (Pat Ellis) — which is what makes it THIS org's projection rather than any populated org's
  97  |     await expect(
  98  |       heroRow.first().getByText(PT_EMPLOYEE.role, { exact: false }).first(),
  99  |       `and her row carries her seeded role ("${PT_EMPLOYEE.role}")`,
  100 |     ).toBeVisible();
  101 |   });
  102 | });
  103 | 
```