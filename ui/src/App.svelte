<script lang="ts">
  import ServerSelector from './ServerSelector.svelte';
  import { serverUrl } from './Stores';
</script>

{#if $serverUrl == ''}
  <ServerSelector />
{:else}
  {$serverUrl}
{/if}

<style>
</style>
