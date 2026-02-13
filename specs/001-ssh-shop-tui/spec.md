# Feature Specification: SSH-Based Shop TUI

**Feature Branch**: `001-ssh-shop-tui`  
**Created**: 2025-02-11  
**Status**: Draft  
**Input**: User description: "shop.gyeongho.dev 는 ssh 기반 버섯 판매 플랫폼의 TUI 어플리케이션이다. 사용자는 ssh -a -i /dev/null shop.gyeongho.dev 명령어를 실행하여 버섯 판매 플랫폼의 TUI 어플리케이션에 접근할 수 있다. 사용자는 TUI에서 단축키로 shop(a) about(s) faq(d) 를 선택해서 각각 버섯 주문, ssh store에 대한 정보, faq 를 확인할 수 있다. 가장 중요한 것은, terminal.shop 프로젝트의 UI/UX를 그대로 따라하여 terminal.shop(커피)의 버섯 판매용 버전을 구현해야 한다는 점이다."

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Access TUI via SSH and Navigate (Priority: P1)

A user runs `ssh -a -i /dev/null shop.gyeongho.dev` and is presented with a terminal user interface (TUI) that matches the look and feel of terminal.shop. The main screen offers clear navigation options: Shop (a), About (s), and FAQ (d). The user can use these single-key shortcuts to move between sections.

**Why this priority**: Without SSH access and a recognizable TUI shell, no other feature is reachable. This is the entry point for the entire experience.

**Independent Test**: Connect via the given SSH command and confirm the TUI appears with navigation labels and shortcuts (a, s, d) visible and responsive.

**Acceptance Scenarios**:

1. **Given** the user has SSH client access, **When** they run `ssh -a -i /dev/null shop.gyeongho.dev`, **Then** a TUI session starts and displays the main navigation (Shop, About, FAQ) with indicated shortcuts.
2. **Given** the TUI is displayed, **When** the user presses the shortcut keys (a, s, d), **Then** the interface switches to the corresponding section (shop, about, faq) or clearly indicates the selection.
3. **Given** the user is in any section, **When** they use the defined navigation or back behavior consistent with terminal.shop, **Then** they can return to the main navigation or move between sections without losing the expected UX.

---

### User Story 2 - Shop: Browse and Order Products (Priority: P2)

A user selects Shop (a) from the main TUI. They can browse products, view details, and complete an order (e.g. add to cart and checkout) using the same interaction patterns as terminal.shop (coffee), adapted for products.

**Why this priority**: The primary business value is selling products; shop is the core transaction flow.

**Independent Test**: Open Shop from the TUI, browse at least one product, and complete a purchase (or simulated checkout) using only the TUI.

**Acceptance Scenarios**:

1. **Given** the user is on the main TUI, **When** they press (a) for Shop, **Then** the shop view opens and lists or allows browsing of products in a terminal.shop-style layout.
2. **Given** the user is in Shop, **When** they select a product and choose to order, **Then** they can add it to a cart (or equivalent) and proceed to a checkout flow.
3. **Given** the user has items in cart, **When** they complete checkout, **Then** they receive clear confirmation and any order summary consistent with terminal.shop UX.

---

### User Story 3 - About: View Store Information (Priority: P3)

A user selects About (s) from the main TUI to learn about the SSH store and the store. They see static content (e.g. story, contact, or branding) presented in the same visual and interaction style as terminal.shop.

**Why this priority**: Trust and context; supports decision-making but is not required to complete a purchase.

**Independent Test**: Open About (s) from the main TUI and verify store information is readable and navigation back to main works.

**Acceptance Scenarios**:

1. **Given** the user is on the main TUI, **When** they press (s) for About, **Then** the about view opens and displays store/SSH store information.
2. **Given** the user is in About, **When** they use the defined back or menu action, **Then** they return to the main navigation.

---

### User Story 4 - FAQ: View Frequently Asked Questions (Priority: P4)

A user selects FAQ (d) from the main TUI to read frequently asked questions and answers. Content is presented in the same TUI style as terminal.shop.

**Why this priority**: Reduces support burden and improves clarity; secondary to completing a purchase.

**Independent Test**: Open FAQ (d) from the main TUI and verify FAQ content is visible and navigable.

**Acceptance Scenarios**:

1. **Given** the user is on the main TUI, **When** they press (d) for FAQ, **Then** the FAQ view opens and displays questions and answers.
2. **Given** the user is in FAQ, **When** they use the defined back or menu action, **Then** they return to the main navigation.

---

### Edge Cases

- What happens when the user presses an unknown key on the main screen? (Show help or ignore; behavior must match terminal.shop-style expectations.)
- How does the system handle SSH disconnection or resize during a session? (Graceful reflow or reconnection guidance.)
- What happens when the user tries to checkout with an empty cart? (Clear message and option to return to shop.)
- How are network or backend errors shown in the TUI? (User-friendly message and way to retry or go back.)

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: System MUST allow access to the TUI via SSH using the command `ssh -a -i /dev/null shop.gyeongho.dev` (or equivalent host/command that reaches the same service).
- **FR-002**: System MUST present a TUI whose UI/UX follows terminal.shop (coffee) so that layout, navigation, shortcuts, and interaction patterns are recognizable as that style.
- **FR-003**: System MUST provide a main navigation with at least three sections: Shop (shortcut a), About (shortcut s), and FAQ (shortcut d).
- **FR-004**: System MUST allow users to open the Shop section and browse products and complete an order (add to cart and checkout) through the TUI.
- **FR-005**: System MUST allow users to open the About section and view store/SSH store information.
- **FR-006**: System MUST allow users to open the FAQ section and view frequently asked questions and answers.
- **FR-007**: System MUST allow users to return from Shop, About, or FAQ to the main navigation using a consistent, terminal.shop-style back or menu action.
- **FR-008**: System MUST display content and prompts in a way that remains usable at common terminal sizes and respects terminal.shop-style layout constraints.

### Key Entities

- **Product**: A sellable product type or SKU; has name, description, and ordering-relevant attributes (e.g. price, availability).
- **Order**: A user’s purchase request; includes selected products, quantities, and fulfillment-relevant data.
- **Store information**: Static content for the About section (e.g. farm story, SSH store description, contact).
- **FAQ entry**: A question-and-answer pair shown in the FAQ section.

## Assumptions

- The hostname `shop.gyeongho.dev` (or the actual SSH host) is configured to accept the specified SSH command and start the TUI session.
- terminal.shop (coffee) is the reference for layout, typography, shortcuts, and flow; implementation will align with that reference.
- Shop flow includes at least: product list/browse, cart (or equivalent), and checkout confirmation.
- About and FAQ content can be static (e.g. config or files) for the initial scope.
- Users have a terminal that supports the TUI (e.g. ANSI, UTF-8, and minimum size); no specific terminal is mandated in the spec.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: A user can connect via the specified SSH command and reach the TUI with visible Shop, About, and FAQ navigation within one session.
- **SC-002**: A user can complete a full shop flow (browse product, add to cart, checkout) using only the TUI in under five minutes.
- **SC-003**: A user familiar with terminal.shop can navigate the shop TUI (main menu, sections, back) without written instructions.
- **SC-004**: About and FAQ sections are reachable via (s) and (d) and display content without errors; users can return to the main menu from both.
