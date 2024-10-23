package wordlegrid

const svgTemplate = `<svg width="348" height="415" version="1.1" viewBox="0 0 348 415" xmlns="http://www.w3.org/2000/svg">
 <rect id="background" width="100%" height="100%" fill="#121213"/>
 <g>
  <rect id="1-1-empty" x="10" y="10" width="60" height="60" display="none" fill="none" stroke="#3a3a3c" stroke-width="2"/>
  <rect id="1-1-absent" x="10" y="10" width="60" height="60" display="none" fill="#3a3a3c" stroke="#3a3a3c" stroke-width="2"/>
  <rect id="1-1-present" x="10" y="10" width="60" height="60" display="none" fill="#b59f3b" stroke="#b59f3b" stroke-width="2"/>
  <rect id="1-1-correct" x="10" y="10" width="60" height="60" display="none" fill="#548c4e" stroke="#548c4e" stroke-width="2"/>
  <rect id="1-2-empty" x="77" y="10" width="60" height="60" display="none" fill="none" stroke="#3a3a3c" stroke-width="2"/>
  <rect id="1-2-absent" x="77" y="10" width="60" height="60" display="none" fill="#3a3a3c" stroke="#3a3a3c" stroke-width="2"/>
  <rect id="1-2-present" x="77" y="10" width="60" height="60" display="none" fill="#b59f3b" stroke="#b59f3b" stroke-width="2"/>
  <rect id="1-2-correct" x="77" y="10" width="60" height="60" display="none" fill="#548c4e" stroke="#548c4e" stroke-width="2"/>
  <rect id="1-3-empty" x="144" y="10" width="60" height="60" display="none" fill="none" stroke="#3a3a3c" stroke-width="2"/>
  <rect id="1-3-absent" x="144" y="10" width="60" height="60" display="none" fill="#3a3a3c" stroke="#3a3a3c" stroke-width="2"/>
  <rect id="1-3-present" x="144" y="10" width="60" height="60" display="none" fill="#b59f3b" stroke="#b59f3b" stroke-width="2"/>
  <rect id="1-3-correct" x="144" y="10" width="60" height="60" display="none" fill="#548c4e" stroke="#548c4e" stroke-width="2"/>
  <rect id="1-4-empty" x="211" y="10" width="60" height="60" display="none" fill="none" stroke="#3a3a3c" stroke-width="2"/>
  <rect id="1-4-absent" x="211" y="10" width="60" height="60" display="none" fill="#3a3a3c" stroke="#3a3a3c" stroke-width="2"/>
  <rect id="1-4-present" x="211" y="10" width="60" height="60" display="none" fill="#b59f3b" stroke="#b59f3b" stroke-width="2"/>
  <rect id="1-4-correct" x="211" y="10" width="60" height="60" display="none" fill="#548c4e" stroke="#548c4e" stroke-width="2"/>
  <rect id="1-5-empty" x="278" y="10" width="60" height="60" display="none" fill="none" stroke="#3a3a3c" stroke-width="2"/>
  <rect id="1-5-absent" x="278" y="10" width="60" height="60" display="none" fill="#3a3a3c" stroke="#3a3a3c" stroke-width="2"/>
  <rect id="1-5-present" x="278" y="10" width="60" height="60" display="none" fill="#b59f3b" stroke="#b59f3b" stroke-width="2"/>
  <rect id="1-5-correct" x="278" y="10" width="60" height="60" display="none" fill="#548c4e" stroke="#548c4e" stroke-width="2"/>
  <text id="1-1-letter" x="40" y="39.7" display="none" dominant-baseline="central" fill="#f8f8f8" font-family="'TN Web Use Only'" font-size="32px" font-weight="bold" text-anchor="middle" xml:space="preserve">$</text>
  <text id="1-2-letter" x="107" y="39.7" display="none" dominant-baseline="central" fill="#f8f8f8" font-family="'TN Web Use Only'" font-size="32px" font-weight="bold" text-anchor="middle" xml:space="preserve">$</text>
  <text id="1-3-letter" x="174" y="39.7" display="none" dominant-baseline="central" fill="#f8f8f8" font-family="'TN Web Use Only'" font-size="32px" font-weight="bold" text-anchor="middle" xml:space="preserve">$</text>
  <text id="1-4-letter" x="241" y="39.7" display="none" dominant-baseline="central" fill="#f8f8f8" font-family="'TN Web Use Only'" font-size="32px" font-weight="bold" text-anchor="middle" xml:space="preserve">$</text>
  <text id="1-5-letter" x="308" y="39.7" display="none" dominant-baseline="central" fill="#f8f8f8" font-family="'TN Web Use Only'" font-size="32px" font-weight="bold" text-anchor="middle" xml:space="preserve">$</text>
  <text id="1-1-letter-absent" x="40" y="39.7" display="none" dominant-baseline="central" fill="#ffffff" font-family="'TN Web Use Only'" font-size="32px" font-weight="bold" text-anchor="middle" xml:space="preserve">$</text>
  <text id="1-2-letter-absent" x="107" y="39.7" display="none" dominant-baseline="central" fill="#ffffff" font-family="'TN Web Use Only'" font-size="32px" font-weight="bold" text-anchor="middle" xml:space="preserve">$</text>
  <text id="1-3-letter-absent" x="174" y="39.7" display="none" dominant-baseline="central" fill="#ffffff" font-family="'TN Web Use Only'" font-size="32px" font-weight="bold" text-anchor="middle" xml:space="preserve">$</text>
  <text id="1-4-letter-absent" x="241" y="39.7" display="none" dominant-baseline="central" fill="#ffffff" font-family="'TN Web Use Only'" font-size="32px" font-weight="bold" text-anchor="middle" xml:space="preserve">$</text>
  <text id="1-5-letter-absent" x="308" y="39.7" display="none" dominant-baseline="central" fill="#ffffff" font-family="'TN Web Use Only'" font-size="32px" font-weight="bold" text-anchor="middle" xml:space="preserve">$</text>
 </g>
 <g transform="translate(0,67)">
  <rect id="2-1-empty" x="10" y="10" width="60" height="60" display="none" fill="none" stroke="#3a3a3c" stroke-width="2"/>
  <rect id="2-1-absent" x="10" y="10" width="60" height="60" display="none" fill="#3a3a3c" stroke="#3a3a3c" stroke-width="2"/>
  <rect id="2-1-present" x="10" y="10" width="60" height="60" display="none" fill="#b59f3b" stroke="#b59f3b" stroke-width="2"/>
  <rect id="2-1-correct" x="10" y="10" width="60" height="60" display="none" fill="#548c4e" stroke="#548c4e" stroke-width="2"/>
  <rect id="2-2-empty" x="77" y="10" width="60" height="60" display="none" fill="none" stroke="#3a3a3c" stroke-width="2"/>
  <rect id="2-2-absent" x="77" y="10" width="60" height="60" display="none" fill="#3a3a3c" stroke="#3a3a3c" stroke-width="2"/>
  <rect id="2-2-present" x="77" y="10" width="60" height="60" display="none" fill="#b59f3b" stroke="#b59f3b" stroke-width="2"/>
  <rect id="2-2-correct" x="77" y="10" width="60" height="60" display="none" fill="#548c4e" stroke="#548c4e" stroke-width="2"/>
  <rect id="2-3-empty" x="144" y="10" width="60" height="60" display="none" fill="none" stroke="#3a3a3c" stroke-width="2"/>
  <rect id="2-3-absent" x="144" y="10" width="60" height="60" display="none" fill="#3a3a3c" stroke="#3a3a3c" stroke-width="2"/>
  <rect id="2-3-present" x="144" y="10" width="60" height="60" display="none" fill="#b59f3b" stroke="#b59f3b" stroke-width="2"/>
  <rect id="2-3-correct" x="144" y="10" width="60" height="60" display="none" fill="#548c4e" stroke="#548c4e" stroke-width="2"/>
  <rect id="2-4-empty" x="211" y="10" width="60" height="60" display="none" fill="none" stroke="#3a3a3c" stroke-width="2"/>
  <rect id="2-4-absent" x="211" y="10" width="60" height="60" display="none" fill="#3a3a3c" stroke="#3a3a3c" stroke-width="2"/>
  <rect id="2-4-present" x="211" y="10" width="60" height="60" display="none" fill="#b59f3b" stroke="#b59f3b" stroke-width="2"/>
  <rect id="2-4-correct" x="211" y="10" width="60" height="60" display="none" fill="#548c4e" stroke="#548c4e" stroke-width="2"/>
  <rect id="2-5-empty" x="278" y="10" width="60" height="60" display="none" fill="none" stroke="#3a3a3c" stroke-width="2"/>
  <rect id="2-5-absent" x="278" y="10" width="60" height="60" display="none" fill="#3a3a3c" stroke="#3a3a3c" stroke-width="2"/>
  <rect id="2-5-present" x="278" y="10" width="60" height="60" display="none" fill="#b59f3b" stroke="#b59f3b" stroke-width="2"/>
  <rect id="2-5-correct" x="278" y="10" width="60" height="60" display="none" fill="#548c4e" stroke="#548c4e" stroke-width="2"/>
  <text id="2-1-letter" x="40" y="39.7" display="none" dominant-baseline="central" fill="#f8f8f8" font-family="'TN Web Use Only'" font-size="32px" font-weight="bold" text-anchor="middle" xml:space="preserve">$</text>
  <text id="2-2-letter" x="107" y="39.7" display="none" dominant-baseline="central" fill="#f8f8f8" font-family="'TN Web Use Only'" font-size="32px" font-weight="bold" text-anchor="middle" xml:space="preserve">$</text>
  <text id="2-3-letter" x="174" y="39.7" display="none" dominant-baseline="central" fill="#f8f8f8" font-family="'TN Web Use Only'" font-size="32px" font-weight="bold" text-anchor="middle" xml:space="preserve">$</text>
  <text id="2-4-letter" x="241" y="39.7" display="none" dominant-baseline="central" fill="#f8f8f8" font-family="'TN Web Use Only'" font-size="32px" font-weight="bold" text-anchor="middle" xml:space="preserve">$</text>
  <text id="2-5-letter" x="308" y="39.7" display="none" dominant-baseline="central" fill="#f8f8f8" font-family="'TN Web Use Only'" font-size="32px" font-weight="bold" text-anchor="middle" xml:space="preserve">$</text>
  <text id="2-1-letter-absent" x="40" y="39.7" display="none" dominant-baseline="central" fill="#ffffff" font-family="'TN Web Use Only'" font-size="32px" font-weight="bold" text-anchor="middle" xml:space="preserve">$</text>
  <text id="2-2-letter-absent" x="107" y="39.7" display="none" dominant-baseline="central" fill="#ffffff" font-family="'TN Web Use Only'" font-size="32px" font-weight="bold" text-anchor="middle" xml:space="preserve">$</text>
  <text id="2-3-letter-absent" x="174" y="39.7" display="none" dominant-baseline="central" fill="#ffffff" font-family="'TN Web Use Only'" font-size="32px" font-weight="bold" text-anchor="middle" xml:space="preserve">$</text>
  <text id="2-4-letter-absent" x="241" y="39.7" display="none" dominant-baseline="central" fill="#ffffff" font-family="'TN Web Use Only'" font-size="32px" font-weight="bold" text-anchor="middle" xml:space="preserve">$</text>
  <text id="2-5-letter-absent" x="308" y="39.7" display="none" dominant-baseline="central" fill="#ffffff" font-family="'TN Web Use Only'" font-size="32px" font-weight="bold" text-anchor="middle" xml:space="preserve">$</text>
 </g>
 <g transform="translate(0,134)">
  <rect id="3-1-empty" x="10" y="10" width="60" height="60" display="none" fill="none" stroke="#3a3a3c" stroke-width="2"/>
  <rect id="3-1-absent" x="10" y="10" width="60" height="60" display="none" fill="#3a3a3c" stroke="#3a3a3c" stroke-width="2"/>
  <rect id="3-1-present" x="10" y="10" width="60" height="60" display="none" fill="#b59f3b" stroke="#b59f3b" stroke-width="2"/>
  <rect id="3-1-correct" x="10" y="10" width="60" height="60" display="none" fill="#548c4e" stroke="#548c4e" stroke-width="2"/>
  <rect id="3-2-empty" x="77" y="10" width="60" height="60" display="none" fill="none" stroke="#3a3a3c" stroke-width="2"/>
  <rect id="3-2-absent" x="77" y="10" width="60" height="60" display="none" fill="#3a3a3c" stroke="#3a3a3c" stroke-width="2"/>
  <rect id="3-2-present" x="77" y="10" width="60" height="60" display="none" fill="#b59f3b" stroke="#b59f3b" stroke-width="2"/>
  <rect id="3-2-correct" x="77" y="10" width="60" height="60" display="none" fill="#548c4e" stroke="#548c4e" stroke-width="2"/>
  <rect id="3-3-empty" x="144" y="10" width="60" height="60" display="none" fill="none" stroke="#3a3a3c" stroke-width="2"/>
  <rect id="3-3-absent" x="144" y="10" width="60" height="60" display="none" fill="#3a3a3c" stroke="#3a3a3c" stroke-width="2"/>
  <rect id="3-3-present" x="144" y="10" width="60" height="60" display="none" fill="#b59f3b" stroke="#b59f3b" stroke-width="2"/>
  <rect id="3-3-correct" x="144" y="10" width="60" height="60" display="none" fill="#548c4e" stroke="#548c4e" stroke-width="2"/>
  <rect id="3-4-empty" x="211" y="10" width="60" height="60" display="none" fill="none" stroke="#3a3a3c" stroke-width="2"/>
  <rect id="3-4-absent" x="211" y="10" width="60" height="60" display="none" fill="#3a3a3c" stroke="#3a3a3c" stroke-width="2"/>
  <rect id="3-4-present" x="211" y="10" width="60" height="60" display="none" fill="#b59f3b" stroke="#b59f3b" stroke-width="2"/>
  <rect id="3-4-correct" x="211" y="10" width="60" height="60" display="none" fill="#548c4e" stroke="#548c4e" stroke-width="2"/>
  <rect id="3-5-empty" x="278" y="10" width="60" height="60" display="none" fill="none" stroke="#3a3a3c" stroke-width="2"/>
  <rect id="3-5-absent" x="278" y="10" width="60" height="60" display="none" fill="#3a3a3c" stroke="#3a3a3c" stroke-width="2"/>
  <rect id="3-5-present" x="278" y="10" width="60" height="60" display="none" fill="#b59f3b" stroke="#b59f3b" stroke-width="2"/>
  <rect id="3-5-correct" x="278" y="10" width="60" height="60" display="none" fill="#548c4e" stroke="#548c4e" stroke-width="2"/>
  <text id="3-1-letter" x="40" y="39.7" display="none" dominant-baseline="central" fill="#f8f8f8" font-family="'TN Web Use Only'" font-size="32px" font-weight="bold" text-anchor="middle" xml:space="preserve">$</text>
  <text id="3-2-letter" x="107" y="39.7" display="none" dominant-baseline="central" fill="#f8f8f8" font-family="'TN Web Use Only'" font-size="32px" font-weight="bold" text-anchor="middle" xml:space="preserve">$</text>
  <text id="3-3-letter" x="174" y="39.7" display="none" dominant-baseline="central" fill="#f8f8f8" font-family="'TN Web Use Only'" font-size="32px" font-weight="bold" text-anchor="middle" xml:space="preserve">$</text>
  <text id="3-4-letter" x="241" y="39.7" display="none" dominant-baseline="central" fill="#f8f8f8" font-family="'TN Web Use Only'" font-size="32px" font-weight="bold" text-anchor="middle" xml:space="preserve">$</text>
  <text id="3-5-letter" x="308" y="39.7" display="none" dominant-baseline="central" fill="#f8f8f8" font-family="'TN Web Use Only'" font-size="32px" font-weight="bold" text-anchor="middle" xml:space="preserve">$</text>
  <text id="3-1-letter-absent" x="40" y="39.7" display="none" dominant-baseline="central" fill="#ffffff" font-family="'TN Web Use Only'" font-size="32px" font-weight="bold" text-anchor="middle" xml:space="preserve">$</text>
  <text id="3-2-letter-absent" x="107" y="39.7" display="none" dominant-baseline="central" fill="#ffffff" font-family="'TN Web Use Only'" font-size="32px" font-weight="bold" text-anchor="middle" xml:space="preserve">$</text>
  <text id="3-3-letter-absent" x="174" y="39.7" display="none" dominant-baseline="central" fill="#ffffff" font-family="'TN Web Use Only'" font-size="32px" font-weight="bold" text-anchor="middle" xml:space="preserve">$</text>
  <text id="3-4-letter-absent" x="241" y="39.7" display="none" dominant-baseline="central" fill="#ffffff" font-family="'TN Web Use Only'" font-size="32px" font-weight="bold" text-anchor="middle" xml:space="preserve">$</text>
  <text id="3-5-letter-absent" x="308" y="39.7" display="none" dominant-baseline="central" fill="#ffffff" font-family="'TN Web Use Only'" font-size="32px" font-weight="bold" text-anchor="middle" xml:space="preserve">$</text>
 </g>
 <g transform="translate(0,201)">
  <rect id="4-1-empty" x="10" y="10" width="60" height="60" display="none" fill="none" stroke="#3a3a3c" stroke-width="2"/>
  <rect id="4-1-absent" x="10" y="10" width="60" height="60" display="none" fill="#3a3a3c" stroke="#3a3a3c" stroke-width="2"/>
  <rect id="4-1-present" x="10" y="10" width="60" height="60" display="none" fill="#b59f3b" stroke="#b59f3b" stroke-width="2"/>
  <rect id="4-1-correct" x="10" y="10" width="60" height="60" display="none" fill="#548c4e" stroke="#548c4e" stroke-width="2"/>
  <rect id="4-2-empty" x="77" y="10" width="60" height="60" display="none" fill="none" stroke="#3a3a3c" stroke-width="2"/>
  <rect id="4-2-absent" x="77" y="10" width="60" height="60" display="none" fill="#3a3a3c" stroke="#3a3a3c" stroke-width="2"/>
  <rect id="4-2-present" x="77" y="10" width="60" height="60" display="none" fill="#b59f3b" stroke="#b59f3b" stroke-width="2"/>
  <rect id="4-2-correct" x="77" y="10" width="60" height="60" display="none" fill="#548c4e" stroke="#548c4e" stroke-width="2"/>
  <rect id="4-3-empty" x="144" y="10" width="60" height="60" display="none" fill="none" stroke="#3a3a3c" stroke-width="2"/>
  <rect id="4-3-absent" x="144" y="10" width="60" height="60" display="none" fill="#3a3a3c" stroke="#3a3a3c" stroke-width="2"/>
  <rect id="4-3-present" x="144" y="10" width="60" height="60" display="none" fill="#b59f3b" stroke="#b59f3b" stroke-width="2"/>
  <rect id="4-3-correct" x="144" y="10" width="60" height="60" display="none" fill="#548c4e" stroke="#548c4e" stroke-width="2"/>
  <rect id="4-4-empty" x="211" y="10" width="60" height="60" display="none" fill="none" stroke="#3a3a3c" stroke-width="2"/>
  <rect id="4-4-absent" x="211" y="10" width="60" height="60" display="none" fill="#3a3a3c" stroke="#3a3a3c" stroke-width="2"/>
  <rect id="4-4-present" x="211" y="10" width="60" height="60" display="none" fill="#b59f3b" stroke="#b59f3b" stroke-width="2"/>
  <rect id="4-4-correct" x="211" y="10" width="60" height="60" display="none" fill="#548c4e" stroke="#548c4e" stroke-width="2"/>
  <rect id="4-5-empty" x="278" y="10" width="60" height="60" display="none" fill="none" stroke="#3a3a3c" stroke-width="2"/>
  <rect id="4-5-absent" x="278" y="10" width="60" height="60" display="none" fill="#3a3a3c" stroke="#3a3a3c" stroke-width="2"/>
  <rect id="4-5-present" x="278" y="10" width="60" height="60" display="none" fill="#b59f3b" stroke="#b59f3b" stroke-width="2"/>
  <rect id="4-5-correct" x="278" y="10" width="60" height="60" display="none" fill="#548c4e" stroke="#548c4e" stroke-width="2"/>
  <text id="4-1-letter" x="40" y="39.7" display="none" dominant-baseline="central" fill="#f8f8f8" font-family="'TN Web Use Only'" font-size="32px" font-weight="bold" text-anchor="middle" xml:space="preserve">$</text>
  <text id="4-2-letter" x="107" y="39.7" display="none" dominant-baseline="central" fill="#f8f8f8" font-family="'TN Web Use Only'" font-size="32px" font-weight="bold" text-anchor="middle" xml:space="preserve">$</text>
  <text id="4-3-letter" x="174" y="39.7" display="none" dominant-baseline="central" fill="#f8f8f8" font-family="'TN Web Use Only'" font-size="32px" font-weight="bold" text-anchor="middle" xml:space="preserve">$</text>
  <text id="4-4-letter" x="241" y="39.7" display="none" dominant-baseline="central" fill="#f8f8f8" font-family="'TN Web Use Only'" font-size="32px" font-weight="bold" text-anchor="middle" xml:space="preserve">$</text>
  <text id="4-5-letter" x="308" y="39.7" display="none" dominant-baseline="central" fill="#f8f8f8" font-family="'TN Web Use Only'" font-size="32px" font-weight="bold" text-anchor="middle" xml:space="preserve">$</text>
  <text id="4-1-letter-absent" x="40" y="39.7" display="none" dominant-baseline="central" fill="#ffffff" font-family="'TN Web Use Only'" font-size="32px" font-weight="bold" text-anchor="middle" xml:space="preserve">$</text>
  <text id="4-2-letter-absent" x="107" y="39.7" display="none" dominant-baseline="central" fill="#ffffff" font-family="'TN Web Use Only'" font-size="32px" font-weight="bold" text-anchor="middle" xml:space="preserve">$</text>
  <text id="4-3-letter-absent" x="174" y="39.7" display="none" dominant-baseline="central" fill="#ffffff" font-family="'TN Web Use Only'" font-size="32px" font-weight="bold" text-anchor="middle" xml:space="preserve">$</text>
  <text id="4-4-letter-absent" x="241" y="39.7" display="none" dominant-baseline="central" fill="#ffffff" font-family="'TN Web Use Only'" font-size="32px" font-weight="bold" text-anchor="middle" xml:space="preserve">$</text>
  <text id="4-5-letter-absent" x="308" y="39.7" display="none" dominant-baseline="central" fill="#ffffff" font-family="'TN Web Use Only'" font-size="32px" font-weight="bold" text-anchor="middle" xml:space="preserve">$</text>
 </g>
 <g transform="translate(0,268)">
  <rect id="5-1-empty" x="10" y="10" width="60" height="60" display="none" fill="none" stroke="#3a3a3c" stroke-width="2"/>
  <rect id="5-1-absent" x="10" y="10" width="60" height="60" display="none" fill="#3a3a3c" stroke="#3a3a3c" stroke-width="2"/>
  <rect id="5-1-present" x="10" y="10" width="60" height="60" display="none" fill="#b59f3b" stroke="#b59f3b" stroke-width="2"/>
  <rect id="5-1-correct" x="10" y="10" width="60" height="60" display="none" fill="#548c4e" stroke="#548c4e" stroke-width="2"/>
  <rect id="5-2-empty" x="77" y="10" width="60" height="60" display="none" fill="none" stroke="#3a3a3c" stroke-width="2"/>
  <rect id="5-2-absent" x="77" y="10" width="60" height="60" display="none" fill="#3a3a3c" stroke="#3a3a3c" stroke-width="2"/>
  <rect id="5-2-present" x="77" y="10" width="60" height="60" display="none" fill="#b59f3b" stroke="#b59f3b" stroke-width="2"/>
  <rect id="5-2-correct" x="77" y="10" width="60" height="60" display="none" fill="#548c4e" stroke="#548c4e" stroke-width="2"/>
  <rect id="5-3-empty" x="144" y="10" width="60" height="60" display="none" fill="none" stroke="#3a3a3c" stroke-width="2"/>
  <rect id="5-3-absent" x="144" y="10" width="60" height="60" display="none" fill="#3a3a3c" stroke="#3a3a3c" stroke-width="2"/>
  <rect id="5-3-present" x="144" y="10" width="60" height="60" display="none" fill="#b59f3b" stroke="#b59f3b" stroke-width="2"/>
  <rect id="5-3-correct" x="144" y="10" width="60" height="60" display="none" fill="#548c4e" stroke="#548c4e" stroke-width="2"/>
  <rect id="5-4-empty" x="211" y="10" width="60" height="60" display="none" fill="none" stroke="#3a3a3c" stroke-width="2"/>
  <rect id="5-4-absent" x="211" y="10" width="60" height="60" display="none" fill="#3a3a3c" stroke="#3a3a3c" stroke-width="2"/>
  <rect id="5-4-present" x="211" y="10" width="60" height="60" display="none" fill="#b59f3b" stroke="#b59f3b" stroke-width="2"/>
  <rect id="5-4-correct" x="211" y="10" width="60" height="60" display="none" fill="#548c4e" stroke="#548c4e" stroke-width="2"/>
  <rect id="5-5-empty" x="278" y="10" width="60" height="60" display="none" fill="none" stroke="#3a3a3c" stroke-width="2"/>
  <rect id="5-5-absent" x="278" y="10" width="60" height="60" display="none" fill="#3a3a3c" stroke="#3a3a3c" stroke-width="2"/>
  <rect id="5-5-present" x="278" y="10" width="60" height="60" display="none" fill="#b59f3b" stroke="#b59f3b" stroke-width="2"/>
  <rect id="5-5-correct" x="278" y="10" width="60" height="60" display="none" fill="#548c4e" stroke="#548c4e" stroke-width="2"/>
  <text id="5-1-letter" x="40" y="39.7" display="none" dominant-baseline="central" fill="#f8f8f8" font-family="'TN Web Use Only'" font-size="32px" font-weight="bold" text-anchor="middle" xml:space="preserve">$</text>
  <text id="5-2-letter" x="107" y="39.7" display="none" dominant-baseline="central" fill="#f8f8f8" font-family="'TN Web Use Only'" font-size="32px" font-weight="bold" text-anchor="middle" xml:space="preserve">$</text>
  <text id="5-3-letter" x="174" y="39.7" display="none" dominant-baseline="central" fill="#f8f8f8" font-family="'TN Web Use Only'" font-size="32px" font-weight="bold" text-anchor="middle" xml:space="preserve">$</text>
  <text id="5-4-letter" x="241" y="39.7" display="none" dominant-baseline="central" fill="#f8f8f8" font-family="'TN Web Use Only'" font-size="32px" font-weight="bold" text-anchor="middle" xml:space="preserve">$</text>
  <text id="5-5-letter" x="308" y="39.7" display="none" dominant-baseline="central" fill="#f8f8f8" font-family="'TN Web Use Only'" font-size="32px" font-weight="bold" text-anchor="middle" xml:space="preserve">$</text>
  <text id="5-1-letter-absent" x="40" y="39.7" display="none" dominant-baseline="central" fill="#ffffff" font-family="'TN Web Use Only'" font-size="32px" font-weight="bold" text-anchor="middle" xml:space="preserve">$</text>
  <text id="5-2-letter-absent" x="107" y="39.7" display="none" dominant-baseline="central" fill="#ffffff" font-family="'TN Web Use Only'" font-size="32px" font-weight="bold" text-anchor="middle" xml:space="preserve">$</text>
  <text id="5-3-letter-absent" x="174" y="39.7" display="none" dominant-baseline="central" fill="#ffffff" font-family="'TN Web Use Only'" font-size="32px" font-weight="bold" text-anchor="middle" xml:space="preserve">$</text>
  <text id="5-4-letter-absent" x="241" y="39.7" display="none" dominant-baseline="central" fill="#ffffff" font-family="'TN Web Use Only'" font-size="32px" font-weight="bold" text-anchor="middle" xml:space="preserve">$</text>
  <text id="5-5-letter-absent" x="308" y="39.7" display="none" dominant-baseline="central" fill="#ffffff" font-family="'TN Web Use Only'" font-size="32px" font-weight="bold" text-anchor="middle" xml:space="preserve">$</text>
 </g>
 <g transform="translate(0,335)">
  <rect id="6-1-empty" x="10" y="10" width="60" height="60" display="none" fill="none" stroke="#3a3a3c" stroke-width="2"/>
  <rect id="6-1-absent" x="10" y="10" width="60" height="60" display="none" fill="#3a3a3c" stroke="#3a3a3c" stroke-width="2"/>
  <rect id="6-1-present" x="10" y="10" width="60" height="60" display="none" fill="#b59f3b" stroke="#b59f3b" stroke-width="2"/>
  <rect id="6-1-correct" x="10" y="10" width="60" height="60" display="none" fill="#548c4e" stroke="#548c4e" stroke-width="2"/>
  <rect id="6-2-empty" x="77" y="10" width="60" height="60" display="none" fill="none" stroke="#3a3a3c" stroke-width="2"/>
  <rect id="6-2-absent" x="77" y="10" width="60" height="60" display="none" fill="#3a3a3c" stroke="#3a3a3c" stroke-width="2"/>
  <rect id="6-2-present" x="77" y="10" width="60" height="60" display="none" fill="#b59f3b" stroke="#b59f3b" stroke-width="2"/>
  <rect id="6-2-correct" x="77" y="10" width="60" height="60" display="none" fill="#548c4e" stroke="#548c4e" stroke-width="2"/>
  <rect id="6-3-empty" x="144" y="10" width="60" height="60" display="none" fill="none" stroke="#3a3a3c" stroke-width="2"/>
  <rect id="6-3-absent" x="144" y="10" width="60" height="60" display="none" fill="#3a3a3c" stroke="#3a3a3c" stroke-width="2"/>
  <rect id="6-3-present" x="144" y="10" width="60" height="60" display="none" fill="#b59f3b" stroke="#b59f3b" stroke-width="2"/>
  <rect id="6-3-correct" x="144" y="10" width="60" height="60" display="none" fill="#548c4e" stroke="#548c4e" stroke-width="2"/>
  <rect id="6-4-empty" x="211" y="10" width="60" height="60" display="none" fill="none" stroke="#3a3a3c" stroke-width="2"/>
  <rect id="6-4-absent" x="211" y="10" width="60" height="60" display="none" fill="#3a3a3c" stroke="#3a3a3c" stroke-width="2"/>
  <rect id="6-4-present" x="211" y="10" width="60" height="60" display="none" fill="#b59f3b" stroke="#b59f3b" stroke-width="2"/>
  <rect id="6-4-correct" x="211" y="10" width="60" height="60" display="none" fill="#548c4e" stroke="#548c4e" stroke-width="2"/>
  <rect id="6-5-empty" x="278" y="10" width="60" height="60" display="none" fill="none" stroke="#3a3a3c" stroke-width="2"/>
  <rect id="6-5-absent" x="278" y="10" width="60" height="60" display="none" fill="#3a3a3c" stroke="#3a3a3c" stroke-width="2"/>
  <rect id="6-5-present" x="278" y="10" width="60" height="60" display="none" fill="#b59f3b" stroke="#b59f3b" stroke-width="2"/>
  <rect id="6-5-correct" x="278" y="10" width="60" height="60" display="none" fill="#548c4e" stroke="#548c4e" stroke-width="2"/>
  <text id="6-1-letter" x="40" y="39.7" display="none" dominant-baseline="central" fill="#f8f8f8" font-family="'TN Web Use Only'" font-size="32px" font-weight="bold" text-anchor="middle" xml:space="preserve">$</text>
  <text id="6-2-letter" x="107" y="39.7" display="none" dominant-baseline="central" fill="#f8f8f8" font-family="'TN Web Use Only'" font-size="32px" font-weight="bold" text-anchor="middle" xml:space="preserve">$</text>
  <text id="6-3-letter" x="174" y="39.7" display="none" dominant-baseline="central" fill="#f8f8f8" font-family="'TN Web Use Only'" font-size="32px" font-weight="bold" text-anchor="middle" xml:space="preserve">$</text>
  <text id="6-4-letter" x="241" y="39.7" display="none" dominant-baseline="central" fill="#f8f8f8" font-family="'TN Web Use Only'" font-size="32px" font-weight="bold" text-anchor="middle" xml:space="preserve">$</text>
  <text id="6-5-letter" x="308" y="39.7" display="none" dominant-baseline="central" fill="#f8f8f8" font-family="'TN Web Use Only'" font-size="32px" font-weight="bold" text-anchor="middle" xml:space="preserve">$</text>
  <text id="6-1-letter-absent" x="40" y="39.7" display="none" dominant-baseline="central" fill="#ffffff" font-family="'TN Web Use Only'" font-size="32px" font-weight="bold" text-anchor="middle" xml:space="preserve">$</text>
  <text id="6-2-letter-absent" x="107" y="39.7" display="none" dominant-baseline="central" fill="#ffffff" font-family="'TN Web Use Only'" font-size="32px" font-weight="bold" text-anchor="middle" xml:space="preserve">$</text>
  <text id="6-3-letter-absent" x="174" y="39.7" display="none" dominant-baseline="central" fill="#ffffff" font-family="'TN Web Use Only'" font-size="32px" font-weight="bold" text-anchor="middle" xml:space="preserve">$</text>
  <text id="6-4-letter-absent" x="241" y="39.7" display="none" dominant-baseline="central" fill="#ffffff" font-family="'TN Web Use Only'" font-size="32px" font-weight="bold" text-anchor="middle" xml:space="preserve">$</text>
  <text id="6-5-letter-absent" x="308" y="39.7" display="none" dominant-baseline="central" fill="#ffffff" font-family="'TN Web Use Only'" font-size="32px" font-weight="bold" text-anchor="middle" xml:space="preserve">$</text>
 </g>
</svg>`
