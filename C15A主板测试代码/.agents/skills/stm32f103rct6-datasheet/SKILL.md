---
name: stm32f103rct6-datasheet
description: "Use when working with the STM32F103RCT6 datasheet, pinout, package pins, GPIO alternate functions, timers, ADC, USART, SPI, I2C, DMA, memory map, boot configuration, clock tree, electrical characteristics, or STM32F103xC/xD/xE hardware limits."
---

# STM32F103RCT6 Datasheet Skill

Use this skill when you need STM32F103RCT6 datasheet facts from the extracted ST PDF instead of relying on memory.

## When To Use This Skill

Use this skill for:

- Pinout and package-pin lookup
- GPIO and alternate-function confirmation
- Timer, PWM, capture/compare, and clock-limit checks
- ADC, DMA, USART, SPI, I2C, CAN, USB, SDIO, FSMC, and RTC capability lookup
- Flash, SRAM, boot mode, reset, power, and memory-map questions
- Electrical characteristics, timing limits, and hardware constraints
- Cross-reference work for STM32F103xC, STM32F103xD, and STM32F103xE

## Primary References

- [Full extracted datasheet](references/STM32F103RCT6.md)
- [Reference index](references/index.md)

## Recommended Workflow

1. Open [references/index.md](references/index.md) to confirm source scope and extraction notes.
2. Search within [references/STM32F103RCT6.md](references/STM32F103RCT6.md) for exact terms such as `GPIO`, `alternate function`, `timer`, `ADC`, `USART`, `SPI`, `I2C`, `DMA`, `memory map`, `electrical characteristics`, `boot`, or exact pin names like `PA9`.
3. Quote the extracted section or nearby page marker when giving hardware guidance.
4. Distinguish clearly between datasheet facts and project-specific implementation assumptions.

## What This Skill Contains

- Full extracted markdown for the 144-page `STM32F103RCT6.PDF`
- Extracted figure/image files under `assets/images/`
- A lightweight reference index for common search terms

## Constraints

- The markdown came from PDF text extraction and contains some encoding artifacts.
- Skill Seekers misclassified many datasheet text fragments as code samples; do not treat those as firmware examples.
- If a line looks suspicious, verify it against surrounding text in the full extracted reference.

## Quick Facts

- Source PDF: `reference_docs/STM32F103RCT6.PDF`
- Source device families in document: `STM32F103xC`, `STM32F103xD`, `STM32F103xE`
- Extracted pages: `144`
- Extracted images preserved by CLI: `17`

## Output Expectations

When using this skill:

- Prefer exact extracted wording over paraphrased memory
- Mark inferences explicitly
- Keep board-specific mappings and control-policy logic outside datasheet facts
