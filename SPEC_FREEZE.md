# 规格冻结表

| 范畴 | 冻结决策 |
|---|---|
|业务边界|语料征集、清洗、许可、标注、分析、发布与撤回；不涉及禁止题材|
|持久化|SQLite 真实 SQL，版本化 schema_migrations，重启复用数据库|
|关系|users/sessions/projects/corpora/batches/annotation_tasks/analyses/audit_events/license_changes，外键与索引|
|事务|配额校验与语料写入同事务；状态变更与审计具备一致性边界|
|状态机|语料、批次、标注任务均拒绝非法转换|
|并发|version 条件更新、唯一索引、单连接 SQLite 事务|
|context|HTTP 到 service、repository、worker 全链路传播并支持取消|
|worker|租约过期重排队、停止信号、周期调度|
|错误传播|稳定领域错误，HTTP 映射 400/401/404/409/500|
|身份权限|登录、8 小时可撤销会话、退出撤销；lead/curator/annotator/reviewer 差异权限|
|HTTP|healthz/readyz、注册登录、项目与语料写入、JSON 响应|
|Docker|真实根入口 go build -o ... .，amd64/arm64 镜像与健康检查|
|测试|领域、service、repository、HTTP、迁移、事务、并发、worker、恢复与错误路径|
|规模|compact_10：生产 Go≥2000 行、≥20 文件、≥10 package；测试 Go≥1500 行|
|容量|冻结至少 10 个独立运行时出题边界，不预埋缺陷、私测或题目分支|
|禁止题材|不做电商、库存、OA、博客、看板、游戏、桌面工具等清单主题|
