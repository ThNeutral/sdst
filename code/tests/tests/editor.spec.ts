import test, { expect, Locator, Page } from "@playwright/test";

test.describe.serial("Editor tests", () => {
  let page1: Page, page2: Page;
  let textarea1: Locator, textarea2: Locator;

  test("Launching first instance", async ({ browser }) => {
    page1 = await browser.newPage();
    await page1.goto("http://localhost:5173/editor");
    await page1.waitForTimeout(600);
    textarea1 = page1.locator(".editor-textarea");
    const text = await textarea1.textContent();
    expect(text?.search("import time")).toBeGreaterThanOrEqual(0);
  });

  test("Launching second instance", async ({ browser }) => {
    page2 = await browser.newPage();
    await page2.goto("http://localhost:5173/editor");
    await page2.waitForTimeout(600);
    textarea2 = page2.locator(".editor-textarea");
    const text = await textarea2.textContent();
    expect(text?.search("import time")).toBeGreaterThanOrEqual(0);
  });

  test("Writing data", async () => {
    const toWrite = "writing data in window1";
    await textarea1.focus();
    await textarea1.press("Control+Home");
    await textarea1.pressSequentially(toWrite);
    await page1.waitForTimeout(300);
    const text2 = await textarea2.textContent();
    expect(text2?.search(toWrite)).toBeGreaterThanOrEqual(0);
  });

  test("Asserting user line marker", async () => {
    const marker = page2.locator('.marker');
    expect(marker).toBeAttached();
    await page2.waitForTimeout(10000);
    const topValue = await marker.evaluate((el) => getComputedStyle(el).top);
    expect(topValue).toBe("40px");
  });

  test("Changing data", async () => {
    const toWrite = "changing data in window1";
    await textarea1.focus();
    await textarea1.press("Control+Home");
    await textarea1.press("Shift+ArrowDown");
    await textarea1.pressSequentially(toWrite);
    await page1.waitForTimeout(300);
    const text2 = await textarea2.textContent();
    expect(text2?.search(toWrite)).toBeGreaterThanOrEqual(0);
  });

  test.afterAll(async () => {
    await textarea1.fill("");
    await textarea1.pressSequentially(
      `import time\n\nwhile True:\n\tprint("test")\n\ttime.sleep(0.5)`
    );
    await page1.waitForTimeout(300);
  });
});
