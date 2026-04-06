# Design System Strategy: Kinetic Precision

## 1. Overview & Creative North Star
The Creative North Star for this design system is **"Kinetic Precision."** 

In a reflex-based environment, the UI must not merely sit on the screen; it must vibrate with the energy of the competition. This system rejects the "templated" look of modern SaaS in favor of a **High-End Digital Brutalism**. We achieve this through aggressive, sharp-edged geometry (0px radii), intentional asymmetry that mimics the speed of the game, and a "Void-to-Neon" contrast ratio. We are moving away from standard grids to create a "Command Center" aesthetic—where every element feels like a high-performance instrument.

The experience is anchored by a deep-space background (`surface`), punctuated by the clashing energies of Player One (`primary` - Electric Blue) and Player Two (`secondary` - Pulse Red).

---

## 2. Colors & Surface Architecture
The color palette is built on high-contrast functionality. We use darkness to provide "depth-of-field," allowing the neon accents to command immediate optical attention.

*   **The "No-Line" Rule:** Under no circumstances are 1px solid borders to be used for sectioning or layout. Boundaries must be defined through **Tonal Shifting**. For example, a card deck area should be defined by placing a `surface-container-low` block against the `background` (#0e0e13). This creates a sophisticated, seamless transition that feels "carved" rather than "drawn."
*   **Surface Hierarchy & Nesting:** Use the surface-container tiers to build a tactile stack. 
    *   **Base:** `surface` (#0e0e13)
    *   **Layout Sections:** `surface-container-low` (#131319)
    *   **Active Interaction Zones:** `surface-container-highest` (#25252d)
*   **The "Glass & Gradient" Rule:** To escape a flat digital feel, use **Vitreous Layering**. For floating HUD elements, utilize `surface-bright` at 40% opacity with a `backdrop-blur` of 20px. 
*   **Signature Textures:** Apply a linear gradient from `primary` (#81ecff) to `primary-container` (#00e3fd) at a 45-degree angle for Electric Blue CTAs. This adds "optical mass" and a premium finish that a flat hex code cannot provide.

---

## 3. Typography: The Action Scale
Typography in this design system is split between **Technical Precision** (Space Grotesk) and **Operational Clarity** (Manrope).

*   **Space Grotesk (Display & Headlines):** This is our "Action" font. It should be used for scores, countdowns, and "KO" states. Use `display-lg` (3.5rem) with tight tracking (-0.02em) to create an editorial, high-impact feel.
*   **Manrope (Body & Titles):** This is the "Strategy" font. It provides the necessary legibility for card descriptions and rules.
*   **Visual Hierarchy:** Leverage extreme scale shifts. A `label-sm` technical stat (0.6875rem) placed next to a `headline-lg` score (2rem) creates the "Editorial Brutalist" tension that defines high-end gaming interfaces.

---

## 4. Elevation & Depth: Tonal Layering
Traditional drop shadows are too "soft" for this system. We achieve depth through light and tone, not blur.

*   **The Layering Principle:** Depth is a result of luminance. Higher-tier containers (e.g., `surface-container-highest`) are perceived as being closer to the user.
*   **Ambient Glows:** Instead of standard black shadows, use **Color-Tinted Ambient Glows**. For a "Pulse Red" element, use a shadow with a 24px blur, 0% offset, and 8% opacity of `secondary` (#ff7073). This makes the UI feel like it is emitting light onto the surface below.
*   **The "Ghost Border" Fallback:** If a container requires further definition, use a "Ghost Border": the `outline-variant` token (#48474d) at 15% opacity. This provides a tactile edge without breaking the "No-Line" rule.
*   **Sharp Corners:** All elements—buttons, cards, and containers—must maintain a **0px corner radius**. This reinforces the "Reflex" theme; speed is sharp, never rounded.

---

## 5. Components

### Buttons: The "Trigger"
*   **Primary (Electric Blue):** Uses the `primary` to `primary-container` gradient. Text is `on-primary` (#005762), Uppercase, Bold Space Grotesk.
*   **Secondary (Pulse Red):** `secondary` (#ff7073) background. Used for high-stakes "Attack" actions.
*   **Interaction:** On hover, the button should not grow; it should emit a `10px` outer glow of its own color.

### Cards: The "Engine"
*   **Structure:** No borders. Use `surface-container-high` for the card body. 
*   **Active State:** When a card is selected, it gains a 2px "Ghost Border" of `tertiary` (#ac89ff) and a subtle `surface-bright` inner glow.
*   **Content:** Separate the "Action Cost" from the "Description" using vertical white space (16px/24px) rather than divider lines.

### Inputs & HUD Elements
*   **Input Fields:** Use `surface-container-lowest` (#000000) for the field background to create a "sunken" feel. The cursor and focus state should use the `primary` Electric Blue.
*   **Chips:** Selection chips should be sharp rectangles. Use `tertiary-container` (#7000ff) with `on-tertiary-container` (#f8f1ff) text to distinguish system-level filters from player-level actions.

### Reflex Indicators
*   **The "Flash" Component:** For millisecond reflex triggers, use a full-screen `surface-variant` overlay at 10% opacity that pulses with the `secondary_dim` (#e90036) color.

---

## 6. Do's and Don'ts

### Do:
*   **Do** use extreme white space to separate card types instead of lines.
*   **Do** lean into asymmetry—offset the player's HUDs slightly to create a sense of kinetic movement.
*   **Do** use `tertiary` (#ac89ff) as a "neutral" energy for system messages (e.g., "Connecting," "Loading").

### Don't:
*   **Don't** ever use a border-radius. Every corner must be a 90-degree angle.
*   **Don't** use grey shadows. If an element needs lift, tint the shadow with the element’s accent color (Blue or Red).
*   **Don't** use standard 1px dividers. If you need a split, use a 4px gap of the `background` color to "slice" the layout.
*   **Don't** settle for centered layouts. Align key stats to the far edges of the container to mimic a high-tech visor or cockpit.