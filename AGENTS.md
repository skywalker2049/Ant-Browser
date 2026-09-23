# Project Agent Instructions

AI Agent 在本仓库工作的基础约束。优先级：用户当轮指令 > 本文 > Agent 默认偏好。

<!-- ant-ready-start-skills:start -->

## Shared Local Skills

Use the shared local skills below when the task matches their scope:

- `page-style-linear-flow`: `D:\code\open_source\ant-ready-start\skills\page-style-linear-flow\SKILL.md`
- `ui-ux-pro-max`: `D:\code\open_source\ant-ready-start\skills\ui-ux-pro-max\SKILL.md`
- `frontend-skill`: `D:\code\open_source\ant-ready-start\skills\frontend-skill\SKILL.md`
- `create-plan`: `D:\code\open_source\ant-ready-start\skills\create-plan\SKILL.md`
- `create-plan-doc`: `D:\code\open_source\ant-ready-start\skills\create-plan-doc\SKILL.md`

Apply it for frontend page design or refactors involving pages, admin panels, forms, tables, dashboards, detail views, wizards, modal/drawer placement, or multi-step flows.

Apply `ui-ux-pro-max` for broader UI/UX design work involving visual direction, design-system shaping, palette and typography selection, component styling, landing pages, dashboards, and cross-stack interface generation when linear-flow rules alone are not enough.

Apply `frontend-skill` when the task needs stronger frontend art direction, visual hierarchy, landing-page composition, sparse premium layouts, image-led sections, or restrained motion design.

Apply `create-plan` when the user explicitly asks for a plan, task breakdown, implementation roadmap, rollout outline, or a step-by-step execution plan before coding.

Apply `create-plan-doc` when the user explicitly asks for a plan that should also be saved into the repository as a markdown document under `docs/plan`.

Core expectations:

- Keep each page focused on one primary responsibility.
- Do not mix operational tables and submit forms on the same screen.
- Use modal/drawer for short low-risk forms; use a dedicated page or wizard for complex flows.
- Remove filler copy, repeated headings, decorative cards, and meaningless whitespace.
- Keep the next action obvious and preserve predictable back/cancel/save behavior.
- 代理连接必须遵守两套连接栈规则，详见 `docs/proxy-connector-stacks.md`：`browser.default_connector_type=xray` 表示 Xray + sing-box 组合栈，Xray 负责 vmess/vless/trojan/shadowsocks/链式代理等，sing-box 负责 hysteria2/tuic/anytls 等协议；`browser.default_connector_type=mihomo` 表示独立 Mihomo 栈。实例启动、测速、真实连通性、IP 健康、预热和代理下载都必须按当前连接栈执行，不允许在 `xray` 组合栈和 `mihomo` 栈之间自动混用；不要把 sing-box 协议误判成“xray 不支持”。
- For detailed UI checks, selectively read `D:\code\open_source\ant-ready-start\skills\page-style-linear-flow\references\checklist.md`.

These shared skill instructions supplement project-specific rules in this `AGENTS.md`; keep more specific project rules authoritative for this repository.

<!-- ant-ready-start-skills:end -->




## 沟通

- 默认中文。先给结论，再给证据，逻辑清晰。不寒暄、不复述、不废话。
- 善用列表、表格、流程图等高效沟通方式。
- 区分「已验证事实」「推断」「未知」，不得把推断写成事实。
- 引用代码给文件路径和行号。
- 用户方案有明显风险时，指出风险 + 给建议方案 + 询问，不静默改需求。

## 调研

- 不猜测未读过的代码或文档。回答或修改前先读直接相关内容和一层调用方。
- 外部 API、版本、兼容性以官方文档和实际接口为准；查不到就说未知，不脑补。
- 能说清要改哪些文件，或连续 2 轮检索无新信息，立即停止检索。够用优于穷尽。

## 决策

- 当前轮消息有明确实现动词才动代码；只问、只查的请求不写实现。授权不跨轮延续。
- 歧义：单一解释直接做；多解释成本相近则选默认并注明假设；工作量差 3 倍以上或影响架构、数据边界、成本，先列选项交用户裁决。
- 不可逆或影响他人的操作（删除、force push、发布、改共享环境）先确认。

## 实现

- MVP 优先，只做当前确认的最小闭环，不为假想需求预建能力。
- 优先复用仓库已有模式，不轻易引入新范式。两个方案都可行时选新增名字、层级、依赖更少的。
- 修根因不修表象。Bugfix 期间不做无关重构，发现的问题作为观察项报告。
- 抽象需两个以上真实重复点支撑。重复优于过早抽象。
- 只在系统边界做校验，不为不可能的场景写防御代码。

## 代码

- 禁止类型压制（`as any`、`@ts-ignore`、`# type: ignore`）绕过错误。
- 禁止空 catch，每个 handler 记录可定位问题的上下文。
- 禁止删除或弱化失败测试来「通过」，禁止硬编码期望值迎合测试。
- 凭据走环境变量或 Secret，不进仓库、日志、Prompt。
- 注释解释「为什么」，风格跟随 codebase。

## 验证

- 无实际执行记录，不得声称「已完成」「已修复」「测试通过」。
- 改完按适用情况全过：诊断无 error；有构建则构建退出码 0；有测试则相关测试通过，已有失败标注为 pre-existing。
- 静态检查只查类型不查逻辑。涉及用户可见行为必须实际跑一遍真实路径。「应该可以」不算验证。
- 每道门通过一次即停，不重复跑已通过的检查。
- 失败如实报告并附输出；没跑就说没跑。

## 失败恢复

- 连续 3 次失败：停止修改 → 回退到最后可用状态 → 记录已尝试方案和精确错误 → 求助或询问用户。
- 重试必须换实质不同的思路，不重复盲试。任何情况下不留下损坏状态。

## Git

- 未经明确要求不 commit、push、rebase、改写历史。
- 不回退、不覆盖用户或其他 Agent 的修改；意料之外的变更视为他人进行中的工作。
- commit 后确认 HEAD 已变更；hook 拒绝先修问题，禁止 `--no-verify`。
- 提交前确认无 secrets、`.env` 被暂存。未经要求不对主干 force push。
- 路径统一用 `/`。空目录不加占位文件。

