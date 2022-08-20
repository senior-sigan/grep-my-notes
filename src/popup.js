document.addEventListener('DOMContentLoaded', async () => {
  const box = document.getElementById('box');
  const tabs = await chrome.tabs.query({ active: true, currentWindow: true });
  const url = new URL(tabs[0].url);
  const params = new URLSearchParams(url.search);

  const query = params.get('q');
  box.innerHTML = query;
  window.query = query;
});
