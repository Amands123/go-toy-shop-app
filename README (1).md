<div align="center">

# 🧸 Go Toy Shop

<img src="https://readme-typing-svg.herokuapp.com?font=Fira+Code&size=22&pause=1000&color=00ADD8&center=true&vCenter=true&width=600&lines=End-to-End+DevSecOps+Project;Build+%7C+Scan+%7C+Deploy+%7C+Monitor;GitHub+Actions+%2B+ArgoCD+%2B+GitOps" alt="Typing SVG" />

</div>

---

<div align="center">

## 🔐 Secure. Automated. Observable. Production-Ready.

</div>

> This project is not just about deploying an application.
> It is about building the **entire delivery system** around it —
> a system where code is automatically **tested**, **quality-checked**, **scanned for vulnerabilities**,
> **containerised**, **deployed without human intervention**, and **monitored in real time**.
>
> Every commit triggers a hardened pipeline that enforces security at three layers —
> source code, container image, and runtime infrastructure —
> before a single pod is updated on the cluster.
>
> This is **DevSecOps** done right.

---
<img width="1366" height="590" alt="image" src="https://github.com/user-attachments/assets/45a26798-244d-4c44-9fd3-d50a0a0cdfa1" />

<div align="center">

[![Go](https://img.shields.io/badge/Go-1.24-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://golang.org)
[![Docker](https://img.shields.io/badge/Docker-Distroless-2496ED?style=for-the-badge&logo=docker&logoColor=white)](https://hub.docker.com/r/hephaestus4i/go-toy-shop)
[![Kubernetes](https://img.shields.io/badge/Kubernetes-GKE-326CE5?style=for-the-badge&logo=kubernetes&logoColor=white)](https://kubernetes.io)
[![Helm](https://img.shields.io/badge/Helm-v3-0F1689?style=for-the-badge&logo=helm&logoColor=white)](https://helm.sh)
[![CI](https://img.shields.io/badge/CI-GitHub_Actions-2088FF?style=for-the-badge&logo=github-actions&logoColor=white)](https://github.com/Amands123/go-toy-shop-app/actions)
[![CD](https://img.shields.io/badge/CD-ArgoCD-EF7B4D?style=for-the-badge&logo=argo&logoColor=white)](https://argo-cd.readthedocs.io)
[![SonarQube](https://img.shields.io/badge/SAST-SonarQube-4E9BCD?style=for-the-badge&logo=sonarqube&logoColor=white)](https://www.sonarsource.com)
[![Trivy](https://img.shields.io/badge/Image_Scan-Trivy-1904DA?style=for-the-badge&logo=aquasecurity&logoColor=white)](https://trivy.dev)
[![Prometheus](https://img.shields.io/badge/Metrics-Prometheus-E6522C?style=for-the-badge&logo=prometheus&logoColor=white)](https://prometheus.io)
[![Grafana](https://img.shields.io/badge/Dashboards-Grafana-F46800?style=for-the-badge&logo=grafana&logoColor=white)](https://grafana.com)

</div>

---

## 📑 Table of Contents

- [Project Overview](#-project-overview)
- [DevSecOps Architecture](#-devsecops-architecture)
- [Project Structure](#-project-structure)
- [Security Layers](#-security-layers)
- [Prerequisites](#-prerequisites)
  - [Install ArgoCD](#-install-argocd)
  - [Install SonarQube](#-install-sonarqube)
  - [Install Prometheus + Grafana](#-install-prometheus--grafana)
- [CI Pipeline — GitHub Actions](#-ci-pipeline--github-actions)
- [CD Pipeline — ArgoCD](#-cd-pipeline--argocd)
- [Monitoring — Prometheus + Grafana](#-monitoring--prometheus--grafana)
- [Local Development](#-local-development)
- [Application Routes](#-application-routes)

---

## 🎯 Project Overview

| Layer | Tool | Purpose |
|---|---|---|
| **Application** | Go 1.24 | Lightweight stateless web app |
| **Containerisation** | Docker + Distroless | Minimal, secure runtime image |
| **CI Pipeline** | GitHub Actions | Automates build, test, scan, push |
| **Code Quality** | SonarQube (self-hosted) | SAST, SCA, coverage, quality gate |
| **Image Security** | Trivy | Scans Docker image for CVEs |
| **Image Registry** | DockerHub | Stores versioned Docker images |
| **CD Pipeline** | ArgoCD | GitOps-based deployment to K8s |
| **Packaging** | Helm v3 | Kubernetes application packaging |
| **Orchestration** | Kubernetes (GKE) | Container runtime & scaling |
| **Monitoring** | Prometheus + Grafana | Metrics collection & dashboards |

---

## 🏗️ DevSecOps Architecture

```mermaid

<h2 class="sr-only">Go Toy Shop — End-to-End DevSecOps Architecture with tool logos</h2>
<div style="font-family:var(--font-sans);padding:12px 4px 0">

<div style="text-align:center;margin-bottom:18px">
  <div style="font-size:10px;font-weight:500;color:var(--color-text-tertiary);letter-spacing:.1em;margin-bottom:3px">END-TO-END DEVSECOPS</div>
  <div style="font-size:19px;font-weight:500;color:var(--color-text-primary)">Go Toy Shop — Architecture</div>
</div>

<!-- ROW 1: SOURCE -->
<div style="display:flex;align-items:center;justify-content:center;gap:10px;margin-bottom:4px">

  <div style="display:flex;align-items:center;gap:8px;padding:9px 14px;border:1px solid var(--color-border-secondary);border-radius:10px;background:var(--color-background-secondary)">
    <svg width="22" height="22" viewBox="0 0 24 24" fill="none" stroke="var(--color-text-secondary)" stroke-width="1.5" stroke-linecap="round"><circle cx="12" cy="8" r="4"/><path d="M6 20v-2a6 6 0 0 1 12 0v2"/></svg>
    <div><div style="font-size:12px;font-weight:500;color:var(--color-text-primary)">Developer</div><div style="font-size:10px;color:var(--color-text-tertiary)">git push</div></div>
  </div>

  <svg width="30" height="14"><line x1="0" y1="7" x2="22" y2="7" stroke="var(--color-border-primary)" stroke-width="1.5"/><path d="M18 3L24 7L18 11" fill="none" stroke="var(--color-border-primary)" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/></svg>

  <div style="display:flex;align-items:center;gap:8px;padding:9px 14px;border:1px solid var(--color-border-secondary);border-radius:10px;background:var(--color-background-secondary)">
    <svg width="22" height="22" viewBox="0 0 24 24" fill="var(--color-text-secondary)"><path d="M12 2C6.477 2 2 6.484 2 12.017c0 4.425 2.865 8.18 6.839 9.504.5.092.682-.217.682-.483 0-.237-.008-.868-.013-1.703-2.782.605-3.369-1.343-3.369-1.343-.454-1.158-1.11-1.466-1.11-1.466-.908-.62.069-.608.069-.608 1.003.07 1.531 1.032 1.531 1.032.892 1.53 2.341 1.088 2.91.832.092-.647.35-1.088.636-1.338-2.22-.253-4.555-1.113-4.555-4.951 0-1.093.39-1.988 1.029-2.688-.103-.253-.446-1.272.098-2.65 0 0 .84-.27 2.75 1.026A9.564 9.564 0 0 1 12 6.844a9.59 9.59 0 0 1 2.504.337c1.909-1.296 2.747-1.027 2.747-1.027.546 1.379.202 2.398.1 2.651.64.7 1.028 1.595 1.028 2.688 0 3.848-2.339 4.695-4.566 4.943.359.309.678.92.678 1.855 0 1.338-.012 2.419-.012 2.747 0 .268.18.58.688.482A10.019 10.019 0 0 0 22 12.017C22 6.484 17.522 2 12 2z"/></svg>
    <div><div style="font-size:12px;font-weight:500;color:var(--color-text-primary)">GitHub</div><div style="font-size:10px;color:var(--color-text-tertiary)">main branch</div></div>
  </div>
</div>

<!-- Arrow down -->
<div style="display:flex;justify-content:center;margin:1px 0 3px">
  <svg width="14" height="22"><line x1="7" y1="0" x2="7" y2="16" stroke="var(--color-border-primary)" stroke-width="1.5"/><path d="M2 11L7 17L12 11" fill="none" stroke="var(--color-border-primary)" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/></svg>
</div>

<!-- CI PIPELINE SECTION -->
<div style="border:1.5px dashed var(--color-border-secondary);border-radius:12px;padding:13px 11px 11px;margin-bottom:4px;position:relative">
  <div style="position:absolute;top:-9px;left:14px;background:var(--color-background-primary);padding:0 7px;font-size:10px;font-weight:500;color:var(--color-text-secondary);letter-spacing:.06em">⚙️ GITHUB ACTIONS — CI PIPELINE</div>

  <div style="display:grid;grid-template-columns:repeat(3,1fr);gap:7px">

    <div style="display:flex;align-items:center;gap:7px;padding:8px 9px;border-radius:8px;background:var(--color-background-secondary);border:1px solid #00ADD830">
      <div style="width:28px;height:28px;border-radius:6px;background:#00ADD8;display:flex;align-items:center;justify-content:center;flex-shrink:0"><span style="font-size:10px;font-weight:700;color:#fff;letter-spacing:-.5px">Go</span></div>
      <div><div style="font-size:11px;font-weight:500;color:var(--color-text-primary)">① Build &amp; Test</div><div style="font-size:9px;color:var(--color-text-tertiary)">go build · go test</div></div>
    </div>

    <div style="display:flex;align-items:center;gap:7px;padding:8px 9px;border-radius:8px;background:var(--color-background-secondary);border:1px solid #4E9BCD30">
      <div style="width:28px;height:28px;border-radius:6px;background:#4E9BCD;display:flex;align-items:center;justify-content:center;flex-shrink:0"><span style="font-size:9px;font-weight:700;color:#fff">SQ</span></div>
      <div><div style="font-size:11px;font-weight:500;color:var(--color-text-primary)">② SonarQube</div><div style="font-size:9px;color:var(--color-text-tertiary)">SAST · SCA · Gate</div></div>
    </div>

    <div style="display:flex;align-items:center;gap:7px;padding:8px 9px;border-radius:8px;background:var(--color-background-secondary);border:1px solid #2496ED30">
      <div style="width:28px;height:28px;border-radius:6px;background:#2496ED;display:flex;align-items:center;justify-content:center;flex-shrink:0">
        <svg width="16" height="16" viewBox="0 0 24 24" fill="white"><path d="M13 4H10V7H13V4ZM13 8H10V11H13V8ZM9 8H6V11H9V8ZM5 8H2V11H5V8ZM9 4H6V7H9V4ZM17 8H14V11H17V8Z"/></svg>
      </div>
      <div><div style="font-size:11px;font-weight:500;color:var(--color-text-primary)">③ Docker Build</div><div style="font-size:9px;color:var(--color-text-tertiary)">distroless · local</div></div>
    </div>

    <div style="display:flex;align-items:center;gap:7px;padding:8px 9px;border-radius:8px;background:var(--color-background-secondary);border:1px solid #1904DA30">
      <div style="width:28px;height:28px;border-radius:6px;background:#1904DA;display:flex;align-items:center;justify-content:center;flex-shrink:0">
        <svg width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="white" stroke-width="2.2" stroke-linecap="round" stroke-linejoin="round"><path d="M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z"/><path d="M9 12l2 2 4-4"/></svg>
      </div>
      <div><div style="font-size:11px;font-weight:500;color:var(--color-text-primary)">④ Trivy Scan</div><div style="font-size:9px;color:var(--color-text-tertiary)">CVE · CRITICAL/HIGH</div></div>
    </div>

    <div style="display:flex;align-items:center;gap:7px;padding:8px 9px;border-radius:8px;background:var(--color-background-secondary);border:1px solid #2496ED30">
      <div style="width:28px;height:28px;border-radius:6px;background:#2496ED;display:flex;align-items:center;justify-content:center;flex-shrink:0">
        <svg width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="white" stroke-width="2.2" stroke-linecap="round"><path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"/><polyline points="17 8 12 3 7 8"/><line x1="12" y1="3" x2="12" y2="15"/></svg>
      </div>
      <div><div style="font-size:11px;font-weight:500;color:var(--color-text-primary)">⑤ Push Image</div><div style="font-size:9px;color:var(--color-text-tertiary)">DockerHub · :run_id</div></div>
    </div>

    <div style="display:flex;align-items:center;gap:7px;padding:8px 9px;border-radius:8px;background:var(--color-background-secondary);border:1px solid #0F168930">
      <div style="width:28px;height:28px;border-radius:6px;background:#0F1689;display:flex;align-items:center;justify-content:center;flex-shrink:0">
        <svg width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="white" stroke-width="2" stroke-linecap="round"><circle cx="12" cy="12" r="3"/><path d="M12 1v4M12 19v4M4.22 4.22l2.83 2.83M16.95 16.95l2.83 2.83M1 12h4M19 12h4M4.22 19.78l2.83-2.83M16.95 7.05l2.83-2.83"/></svg>
      </div>
      <div><div style="font-size:11px;font-weight:500;color:var(--color-text-primary)">⑥ Update Helm</div><div style="font-size:9px;color:var(--color-text-tertiary)">auto-commit tag</div></div>
    </div>

  </div>
</div>

<!-- 3 arrows down -->
<div style="display:grid;grid-template-columns:1fr 1fr 1fr;gap:7px;margin:1px 0 4px">
  <div style="display:flex;flex-direction:column;align-items:center"><div style="font-size:8px;color:var(--color-text-tertiary)">quality gate</div><svg width="14" height="18"><line x1="7" y1="0" x2="7" y2="12" stroke="var(--color-border-primary)" stroke-width="1.5"/><path d="M2 8L7 14L12 8" fill="none" stroke="var(--color-border-primary)" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/></svg></div>
  <div style="display:flex;flex-direction:column;align-items:center"><div style="font-size:8px;color:var(--color-text-tertiary)">push image</div><svg width="14" height="18"><line x1="7" y1="0" x2="7" y2="12" stroke="var(--color-border-primary)" stroke-width="1.5"/><path d="M2 8L7 14L12 8" fill="none" stroke="var(--color-border-primary)" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/></svg></div>
  <div style="display:flex;flex-direction:column;align-items:center"><div style="font-size:8px;color:var(--color-text-tertiary)">git commit</div><svg width="14" height="18"><line x1="7" y1="0" x2="7" y2="12" stroke="var(--color-border-primary)" stroke-width="1.5"/><path d="M2 8L7 14L12 8" fill="none" stroke="var(--color-border-primary)" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/></svg></div>
</div>

<!-- EXTERNAL SERVICES ROW -->
<div style="display:grid;grid-template-columns:1fr 1fr 1fr;gap:7px;margin-bottom:4px">

  <div style="display:flex;align-items:center;gap:8px;padding:9px 11px;border:1px solid #4E9BCD40;border-radius:10px;background:#4E9BCD0D">
    <div style="width:28px;height:28px;border-radius:7px;background:#4E9BCD;display:flex;align-items:center;justify-content:center;flex-shrink:0"><span style="font-size:10px;font-weight:700;color:#fff">SQ</span></div>
    <div><div style="font-size:11px;font-weight:500;color:var(--color-text-primary)">SonarQube</div><div style="font-size:9px;color:var(--color-text-tertiary)">self-hosted :9000</div></div>
  </div>

  <div style="display:flex;align-items:center;gap:8px;padding:9px 11px;border:1px solid #2496ED40;border-radius:10px;background:#2496ED0D">
    <div style="width:28px;height:28px;border-radius:7px;background:#2496ED;display:flex;align-items:center;justify-content:center;flex-shrink:0">
      <svg width="16" height="16" viewBox="0 0 24 24" fill="white"><path d="M13 4H10V7H13V4ZM9 8H6V11H9V8ZM13 8H10V11H13V8ZM17 8H14V11H17V8ZM5 8H2V11H5V8ZM9 4H6V7H9V4Z"/></svg>
    </div>
    <div><div style="font-size:11px;font-weight:500;color:var(--color-text-primary)">DockerHub</div><div style="font-size:9px;color:var(--color-text-tertiary)">hephaestus4i</div></div>
  </div>

  <div style="display:flex;align-items:center;gap:8px;padding:9px 11px;border:1px solid #EF7B4D40;border-radius:10px;background:#EF7B4D0D">
    <div style="width:28px;height:28px;border-radius:7px;background:#EF7B4D;display:flex;align-items:center;justify-content:center;flex-shrink:0">
      <svg width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="white" stroke-width="2" stroke-linecap="round"><path d="M21 16V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16z"/></svg>
    </div>
    <div><div style="font-size:11px;font-weight:500;color:var(--color-text-primary)">ArgoCD detects</div><div style="font-size:9px;color:var(--color-text-tertiary)">values.yaml commit</div></div>
  </div>
</div>

<!-- Arrow down to K8s -->
<div style="display:flex;justify-content:center;margin:1px 0 4px">
  <div style="display:flex;flex-direction:column;align-items:center">
    <div style="font-size:8px;color:var(--color-text-tertiary)">helm sync · image pull</div>
    <svg width="14" height="18"><line x1="7" y1="0" x2="7" y2="12" stroke="var(--color-border-primary)" stroke-width="1.5"/><path d="M2 8L7 14L12 8" fill="none" stroke="var(--color-border-primary)" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/></svg>
  </div>
</div>

<!-- KUBERNETES CLUSTER -->
<div style="border:1.5px solid var(--color-border-secondary);border-radius:12px;padding:13px 11px 11px;margin-bottom:8px;position:relative">
  <div style="position:absolute;top:-9px;left:14px;background:var(--color-background-primary);padding:0 7px;font-size:10px;font-weight:500;color:var(--color-text-secondary);letter-spacing:.06em">☸️ KUBERNETES CLUSTER — GKE</div>

  <!-- Platform tools -->
  <div style="font-size:9px;font-weight:500;color:var(--color-text-tertiary);margin-bottom:6px;letter-spacing:.05em">PLATFORM TOOLS</div>
  <div style="display:grid;grid-template-columns:repeat(4,1fr);gap:6px;margin-bottom:10px">

    <div style="display:flex;flex-direction:column;align-items:center;gap:4px;padding:8px 4px;border-radius:8px;background:var(--color-background-secondary);border:1px solid #EF7B4D30">
      <div style="width:26px;height:26px;border-radius:6px;background:#EF7B4D;display:flex;align-items:center;justify-content:center">
        <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="white" stroke-width="2" stroke-linecap="round"><path d="M21 16V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16z"/></svg>
      </div>
      <div style="font-size:10px;font-weight:500;color:var(--color-text-primary)">ArgoCD</div>
      <div style="font-size:8px;color:var(--color-text-tertiary)">GitOps CD</div>
    </div>

    <div style="display:flex;flex-direction:column;align-items:center;gap:4px;padding:8px 4px;border-radius:8px;background:var(--color-background-secondary);border:1px solid #4E9BCD30">
      <div style="width:26px;height:26px;border-radius:6px;background:#4E9BCD;display:flex;align-items:center;justify-content:center"><span style="font-size:10px;font-weight:700;color:#fff">SQ</span></div>
      <div style="font-size:10px;font-weight:500;color:var(--color-text-primary)">SonarQube</div>
      <div style="font-size:8px;color:var(--color-text-tertiary)">:9000</div>
    </div>

    <div style="display:flex;flex-direction:column;align-items:center;gap:4px;padding:8px 4px;border-radius:8px;background:var(--color-background-secondary);border:1px solid #E6522C30">
      <div style="width:26px;height:26px;border-radius:6px;background:#E6522C;display:flex;align-items:center;justify-content:center">
        <svg width="13" height="13" viewBox="0 0 24 24" fill="white"><circle cx="12" cy="12" r="10" fill="white" fill-opacity="0.2"/><circle cx="12" cy="12" r="5" fill="white" fill-opacity="0.4"/><circle cx="12" cy="12" r="2" fill="white"/></svg>
      </div>
      <div style="font-size:10px;font-weight:500;color:var(--color-text-primary)">Prometheus</div>
      <div style="font-size:8px;color:var(--color-text-tertiary)">metrics</div>
    </div>

    <div style="display:flex;flex-direction:column;align-items:center;gap:4px;padding:8px 4px;border-radius:8px;background:var(--color-background-secondary);border:1px solid #F4680030">
      <div style="width:26px;height:26px;border-radius:6px;background:#F46800;display:flex;align-items:center;justify-content:center"><span style="font-size:12px;font-weight:700;color:#fff">G</span></div>
      <div style="font-size:10px;font-weight:500;color:var(--color-text-primary)">Grafana</div>
      <div style="font-size:8px;color:var(--color-text-tertiary)">dashboards</div>
    </div>
  </div>

  <!-- Divider -->
  <div style="border-top:1px dashed var(--color-border-tertiary);margin-bottom:10px"></div>

  <!-- App -->
  <div style="font-size:9px;font-weight:500;color:var(--color-text-tertiary);margin-bottom:6px;letter-spacing:.05em">GO TOY SHOP APPLICATION</div>
  <div style="display:flex;align-items:center;gap:6px">

    <div style="flex:0 0 auto;display:flex;flex-direction:column;align-items:center;gap:3px;padding:7px 10px;border-radius:8px;background:var(--color-background-secondary);border:1px solid var(--color-border-tertiary)">
      <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="var(--color-text-secondary)" stroke-width="1.5" stroke-linecap="round"><circle cx="12" cy="12" r="10"/><line x1="2" y1="12" x2="22" y2="12"/><path d="M12 2a15.3 15.3 0 0 1 4 10 15.3 15.3 0 0 1-4 10 15.3 15.3 0 0 1-4-10 15.3 15.3 0 0 1 4-10z"/></svg>
      <div style="font-size:9px;font-weight:500;color:var(--color-text-primary)">Ingress</div>
    </div>

    <svg width="14" height="12" style="flex-shrink:0"><line x1="0" y1="6" x2="8" y2="6" stroke="var(--color-border-primary)" stroke-width="1.5"/><path d="M6 2L10 6L6 10" fill="none" stroke="var(--color-border-primary)" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/></svg>

    <div style="flex:0 0 auto;display:flex;flex-direction:column;align-items:center;gap:3px;padding:7px 10px;border-radius:8px;background:var(--color-background-secondary);border:1px solid var(--color-border-tertiary)">
      <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="var(--color-text-secondary)" stroke-width="1.5" stroke-linecap="round"><polyline points="16 3 21 3 21 8"/><line x1="4" y1="20" x2="21" y2="3"/><polyline points="21 16 21 21 16 21"/><line x1="15" y1="15" x2="21" y2="21"/></svg>
      <div style="font-size:9px;font-weight:500;color:var(--color-text-primary)">Service</div>
      <div style="font-size:8px;color:var(--color-text-tertiary)">ClusterIP</div>
    </div>

    <svg width="14" height="12" style="flex-shrink:0"><line x1="0" y1="6" x2="8" y2="6" stroke="var(--color-border-primary)" stroke-width="1.5"/><path d="M6 2L10 6L6 10" fill="none" stroke="var(--color-border-primary)" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/></svg>

    <div style="flex:1;display:grid;grid-template-columns:1fr 1fr;gap:6px">
      <div style="display:flex;flex-direction:column;align-items:center;gap:3px;padding:7px 6px;border-radius:8px;background:var(--color-background-secondary);border:1px solid #00ADD840">
        <div style="width:20px;height:20px;border-radius:5px;background:#00ADD8;display:flex;align-items:center;justify-content:center"><span style="font-size:8px;font-weight:700;color:#fff">Go</span></div>
        <div style="font-size:9px;font-weight:500;color:var(--color-text-primary)">Pod 1</div>
        <div style="font-size:8px;color:var(--color-text-tertiary)">:8080</div>
      </div>
      <div style="display:flex;flex-direction:column;align-items:center;gap:3px;padding:7px 6px;border-radius:8px;background:var(--color-background-secondary);border:1px solid #00ADD840">
        <div style="width:20px;height:20px;border-radius:5px;background:#00ADD8;display:flex;align-items:center;justify-content:center"><span style="font-size:8px;font-weight:700;color:#fff">Go</span></div>
        <div style="font-size:9px;font-weight:500;color:var(--color-text-primary)">Pod 2</div>
        <div style="font-size:8px;color:var(--color-text-tertiary)">:8080</div>
      </div>
    </div>

  </div>
</div>

<!-- User at bottom -->
<div style="display:flex;justify-content:center;align-items:center;gap:8px">
  <svg width="14" height="18" style="transform:rotate(180deg)"><line x1="7" y1="0" x2="7" y2="12" stroke="var(--color-border-primary)" stroke-width="1.5"/><path d="M2 8L7 14L12 8" fill="none" stroke="var(--color-border-primary)" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/></svg>
  <div style="display:flex;align-items:center;gap:8px;padding:8px 16px;border:1px solid var(--color-border-secondary);border-radius:10px;background:var(--color-background-secondary)">
    <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="var(--color-text-secondary)" stroke-width="1.5" stroke-linecap="round"><circle cx="12" cy="12" r="10"/><line x1="2" y1="12" x2="22" y2="12"/><path d="M12 2a15.3 15.3 0 0 1 4 10 15.3 15.3 0 0 1-4 10 15.3 15.3 0 0 1-4-10 15.3 15.3 0 0 1 4-10z"/></svg>
    <div style="font-size:12px;font-weight:500;color:var(--color-text-primary)">Users — Browser</div>
  </div>
</div>

</div>


flowchart TD
    DEV(["👨‍💻 Developer\ngit push to main"])
    GH(["🐙 GitHub Repository"])
    DEV -->|git push| GH

    subgraph CI ["⚙️ GitHub Actions — CI Pipeline"]
        direction TB
        A["① Build & Test\ngo build · go test\ncoverage.out generated"]
        B["② SonarQube Scan\nSAST · SCA · Quality Gate\nBlocks pipeline if gate fails"]
        C["③ Build Docker Image\nMulti-stage · Distroless\nLoaded locally for scanning"]
        D["④ Trivy Image Scan\nCVE scan · CRITICAL & HIGH\nReport uploaded as artifact"]
        E["⑤ Push to DockerHub\nTagged with github.run_id\nImmutable · Traceable"]
        F["⑥ Update Helm Chart\nNew image tag → values.yaml\nAuto-committed to Git"]
        A --> B --> C --> D --> E --> F
    end

    GH --> CI

    SQ(["🔵 SonarQube\nSelf-hosted on GKE\nport 9000"])
    DH(["🐋 DockerHub\nhephaestus4i/go-toy-shop\n:run_id"])

    B <-->|scan + quality gate| SQ
    E -->|push image| DH

    subgraph K8S ["☸️ Kubernetes Cluster — GKE"]
        direction TB

        subgraph TOOLS ["🛠️ Platform Tools"]
            ARGO["🟠 ArgoCD\nGitOps CD · port 80"]
            PROM["🔴 Prometheus\nMetrics scraping · port 9090"]
            GRAF["🟡 Grafana\nDashboards · port 80"]
            SQK["🔵 SonarQube\nCode quality · port 9000"]
            PROM -->|visualise| GRAF
        end

        subgraph APP ["🧸 Go Toy Shop"]
            ING["🌐 Ingress\n/ → service:80"]
            SVC["🔀 Service ClusterIP\nport 80 → 8080"]
            DEP["📋 Deployment · 2 Replicas"]
            P1["🟢 Pod 1\ngo-toy-shop:run_id"]
            P2["🟢 Pod 2\ngo-toy-shop:run_id"]
            ING --> SVC --> DEP --> P1 & P2
        end

        PROM -->|scrape metrics| APP
    end

    F -->|commit detected| ARGO
    ARGO -->|helm sync| DEP
    DH -->|image pull| P1 & P2

    USER(["🌍 User"]) -->|HTTP| ING
```

---

## 📁 Project Structure

```
go-toy-shop-app/
│
├── main.go                          # HTTP server & route registration
├── main_test.go                     # Handler tests
├── go.mod                           # Go module (v1.24)
│
├── handlers/
│   └── handlers.go                  # Request handlers & in-memory data
│
├── templates/                       # HTML templates
│   ├── home.html
│   ├── toys.html
│   ├── cart.html
│   ├── order-success.html
│   ├── about.html
│   └── contact.html
│
├── static/                          # CSS & images
│
├── Dockerfile                       # Multi-stage: golang:1.24 → distroless
├── sonar-project.properties         # SonarQube project config
│
├── helm/
│   └── go-toy-shop-chart/
│       ├── Chart.yaml
│       ├── values.yaml              # ← image.tag auto-updated by CI
│       └── templates/
│           ├── deployment.yaml
│           ├── service.yaml
│           └── ingress.yaml
│
├── k8s/manifests/                   # Raw Kubernetes manifests
│
└── .github/
    └── workflows/
        └── ci.yaml                  # GitHub Actions CI pipeline
```

---

## 🔐 Security Layers

```
Layer 1 — Source Code        SonarQube SAST/SCA
                              Scans Go code for vulnerabilities,
                              code smells, and security hotspots.
                              Quality Gate must PASS to proceed.
                                      ↓
Layer 2 — Container Image    Trivy CVE Scanner
                              Scans built Docker image for
                              OS and dependency vulnerabilities.
                              Report saved as pipeline artifact.
                                      ↓
Layer 3 — Runtime            Distroless Base Image
                              No shell. No package manager.
                              Minimal attack surface in production.
                                      ↓
Layer 4 — Config Drift       ArgoCD Self-Heal
                              Any manual cluster change is
                              automatically reverted to match Git.
```

---

## ✅ Prerequisites

| Requirement | Details |
|---|---|
| GKE Cluster | Must be running with `kubectl` configured |
| Helm v3 | Required for all tool installations |
| ArgoCD | Installed on cluster — see steps below |
| SonarQube | Installed on cluster — see steps below |
| Prometheus + Grafana | Installed on cluster — see steps below |

---

## 🟠 Install ArgoCD

```bash
# 1. Create namespace
kubectl create namespace argocd

# 2. Install ArgoCD
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml

# 3. Expose UI externally
kubectl patch svc argocd-server -n argocd -p '{"spec": {"type": "LoadBalancer"}}'

# 4. Get external IP
kubectl get svc argocd-server -n argocd

# 5. Get admin password
kubectl -n argocd get secret argocd-initial-admin-secret \
  -o jsonpath="{.data.password}" | base64 -d; echo
```

> 🌐 `http://<EXTERNAL-IP>` &nbsp;|&nbsp; Login: `admin` / `<password from step 5>`

### Connect ArgoCD to your repo

```yaml
# argocd-app.yaml
apiVersion: argoproj.io/v1alpha1
kind: Application
metadata:
  name: go-toy-shop
  namespace: argocd
spec:
  project: default
  source:
    repoURL: https://github.com/Amands123/go-toy-shop-app.git
    targetRevision: main
    path: helm/go-toy-shop-chart
  destination:
    server: https://kubernetes.default.svc
    namespace: default
  syncPolicy:
    automated:
      prune: true
      selfHeal: true
```

```bash
kubectl apply -f argocd-app.yaml
```

---

## 🔵 Install SonarQube

```bash
# 1. Create namespace
kubectl create namespace sonarqube

# 2. Add Helm repo
helm repo add sonarqube https://SonarSource.github.io/helm-chart-sonarqube
helm repo update

# 3. Install SonarQube Community Edition
helm install sonarqube sonarqube/sonarqube -n sonarqube \
  --set monitoringPasscode="admin123" \
  --set community.enabled=true

# 4. Wait until Running  (takes 2-3 mins due to Elasticsearch)
kubectl get pods -n sonarqube -w

# 5. Expose UI externally
kubectl patch svc sonarqube-sonarqube -n sonarqube -p '{"spec": {"type": "LoadBalancer"}}'

# 6. Get external IP
kubectl get svc sonarqube-sonarqube -n sonarqube
```

> 🌐 `http://<EXTERNAL-IP>:9000` &nbsp;|&nbsp; Login: `admin` / `admin`

### Generate token for CI

```
Login → My Account → Security → Generate Token → copy value
```

Add `sonar-project.properties` at repo root:

```properties
sonar.projectKey=go-toy-shop-app
sonar.projectName=Go Toy Shop
sonar.sources=.
sonar.exclusions=**/vendor/**,**/*_test.go
sonar.tests=.
sonar.test.inclusions=**/*_test.go
sonar.go.coverage.reportPaths=coverage.out
sonar.language=go
```

---

## 📊 Install Prometheus + Grafana

```bash
# 1. Create namespace
kubectl create namespace monitoring

# 2. Add Helm repo
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update

# 3. Install kube-prometheus-stack (Prometheus + Grafana + Alertmanager)
helm install kube-prometheus-stack prometheus-community/kube-prometheus-stack \
  --namespace monitoring \
  --set grafana.adminPassword="admin123" \
  --set grafana.service.type=LoadBalancer \
  --set prometheus.prometheusSpec.retention=15d

# 4. Watch pods come up
kubectl get pods -n monitoring -w

# 5. Get Grafana external IP
kubectl get svc -n monitoring | grep grafana

# 6. Expose Prometheus externally (optional)
kubectl patch svc kube-prometheus-stack-prometheus -n monitoring \
  -p '{"spec": {"type": "LoadBalancer"}}'
kubectl get svc kube-prometheus-stack-prometheus -n monitoring
```

> 🌐 Grafana: `http://<EXTERNAL-IP>` &nbsp;|&nbsp; Login: `admin` / `admin123`
> 🌐 Prometheus: `http://<EXTERNAL-IP>:9090`

### Pre-built Grafana dashboards to import

```
Dashboards → Import → Enter ID → Load
```

| Dashboard | ID | Shows |
|---|---|---|
| Kubernetes Cluster | `7249` | Node CPU, memory, disk |
| Kubernetes Pods | `6417` | Per-pod resource usage |
| Kubernetes Deployments | `8588` | Deployment health & restarts |
| Go Application | `10826` | Go runtime metrics |

### Enable Prometheus scraping for your app

Add to `helm/go-toy-shop-chart/templates/deployment.yaml`:

```yaml
spec:
  template:
    metadata:
      annotations:
        prometheus.io/scrape: "true"
        prometheus.io/port: "8080"
        prometheus.io/path: "/metrics"
```

---

## ⚙️ CI Pipeline — GitHub Actions

> Triggered on every push to `main` — ignores `helm/`, `k8s/`, `README.md`

```
push to main
    │
    ▼
① build              go mod tidy → go build → go test → upload coverage.out
    │
    ▼
② code-quality       SonarQube SAST/SCA scan + Quality Gate
    │                 ✗ Gate fails → pipeline stops, nothing ships
    ▼
③ push               Build Docker image locally (push: false, load: true)
    │                 Trivy scans local image for CVEs
    │                 Report uploaded as artifact
    │                 Push clean image to DockerHub with :run_id tag
    ▼
④ update-helm        sed new image tag into helm/values.yaml
                      git commit & push → triggers ArgoCD sync
```

### Job dependency chain

| Job | Needs | Security gate |
|---|---|---|
| `build` | — | Tests must pass |
| `code-quality` | `build` | SonarQube quality gate must pass |
| `push` | `code-quality` | Trivy scan runs before push |
| `update-newtag-in-helm-chart` | `push` | Image must be scanned and pushed |

### Required GitHub Secrets

| Secret | Where to get it |
|---|---|
| `SONAR_TOKEN` | SonarQube → My Account → Security |
| `SONAR_HOST_URL` | `http://<sonarqube-external-ip>:9000` |
| `DOCKERHUB_USERNAME` | Your DockerHub username |
| `DOCKERHUB_TOKEN` | DockerHub → Account Settings → Security |
| `TOKEN` | GitHub → Settings → Developer Settings → PAT |

---

## 🔄 CD Pipeline — ArgoCD

> ArgoCD is the **only** CD mechanism. No `helm upgrade` in CI.

```
CI commits new image tag to helm/values.yaml
            │
            ▼
ArgoCD detects commit  (polls every 3 mins)
            │
            ▼
Runs helm sync → rolling update on cluster
            │
            ▼
New pods pull updated image from DockerHub
Old pods terminate after health checks pass
            │
            ▼
selfHeal: true → any manual cluster change
is automatically reverted to match Git
```

### Monitor deployments

```bash
kubectl get applications -n argocd
kubectl get pods -w
kubectl describe pod -l app=go-toy-shop | grep Image
```

---

## 📈 Monitoring — Prometheus + Grafana

```
Prometheus scrapes metrics from:
  ├── go-toy-shop pods      (app metrics on :8080/metrics)
  ├── Node Exporter         (CPU, memory, disk per node)
  ├── kube-state-metrics    (pod restarts, deployment health)
  └── Kubernetes components (API server, scheduler, etcd)
            │
            ▼
Grafana visualises everything via dashboards
            │
            ▼
Alertmanager fires alerts when thresholds are breached
```

---

## 💻 Local Development

```bash
# Clone
git clone https://github.com/Amands123/go-toy-shop-app.git
cd go-toy-shop-app

# Run
go mod tidy
go run main.go
# → http://localhost:8080

# Test
go test ./... -v
go test ./... -cover

# Docker
docker build -t go-toy-shop:local .
docker run -p 8080:8080 go-toy-shop:local
```

---

## 🛣️ Application Routes

| Method | Route | Description |
|---|---|---|
| `GET` | `/` | Home — toy categories |
| `GET` | `/toys?category=<name>` | Products by category |
| `GET` | `/add-to-cart?id=<id>` | Add item to cart |
| `GET` | `/cart` | View cart |
| `POST` | `/place-order` | Place order |
| `GET` | `/about` | About page |
| `GET` | `/contact` | Contact page |

---

<div align="center">

### 🔐 Build it. Secure it. Ship it. Watch it.

**Made with ❤️ by [Aman Agarwal](https://github.com/Amands123)**

[![GitHub](https://img.shields.io/badge/GitHub-Amands123-181717?style=flat-square&logo=github)](https://github.com/Amands123)
[![DockerHub](https://img.shields.io/badge/DockerHub-hephaestus4i-2496ED?style=flat-square&logo=docker&logoColor=white)](https://hub.docker.com/u/hephaestus4i)

</div>
