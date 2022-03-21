<template>
  <div class="brief-item">
    <h1 class="brief-title">
      Systrace 线程 CPU 运行状态分析技巧 - Sleep 和 Uninterruptible Sleep 篇
    </h1>
    <h4 class="brief-content">
      本文是 Systrace 线程 CPU 运行状态分析技巧系列的第三篇，本文主要讲了使用
      Systrace 分析 CPU 状态时遇到的 Sleep 与 Uninterruptible Sleep
      状态的原因排查方法与优化方法，这两个状态导致性能变差概率非常高，而且排查起来也比较费劲，网上也没有系统化的文档。
      本系列的目的是通过 Systrace 这个工具，从另外一个角度来看待 Android
      系统整体的运行，同时也从另外一个角度来对 Framework
      进行学习。也许你看了很多讲 Framework
      的文章，但是总是记不住代码，或者不清楚其运行的流程，也许从 Systrace
      这个图形化的角度，你可以理解的更深入一些。Systrace
      基础和实战系列大家可以在 Systrace 基础知识 - Systrace 预备知识 或者
      博客文章目录 这里看到完整的目录
    </h4>
    <div class="brief-footer">
      <div class="brief-date-content">2022/03/13</div>
      <div class="brief-tags">Android Performance Systrace</div>
    </div>
  </div>
</template>