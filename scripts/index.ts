import { $ } from "bun";
import { array, object, string } from "yup";

const getLatestRelease = async () => {
	console.log("Fetching latest release...");
	const AssetSchema = object({
		name: string().required(),
		browser_download_url: string().required(),
	});
	const ReleaseSchema = object({
		name: string().required(),
		assets: array().of(AssetSchema.required()).required(),
	});
	const json = await fetch(
		"https://api.github.com/repos/mayswind/AriaNg/releases/latest",
	)
		.then((res) => res.json())
		.then((data) => ReleaseSchema.validate(data, { stripUnknown: true }));
	return json;
};
const getDownloadUrl = (targetAsset: string) => {
	console.log("Retrieving download URL...");
	const allinoneAsset = json.assets.find((x) => x.name.includes(targetAsset));
	if (!allinoneAsset)
		throw new Error(`${targetAsset} asset not found in ${json.name}`);
	console.log(`Downloading ${json.name} ${targetAsset}...`);
	const downloadUrl = allinoneAsset.browser_download_url;
	return downloadUrl;
};
const downloadAndExtract = async (targetAsset: string, downloadUrl: string) => {
	console.log("Downloading and extracting...");
	await $`mkdir -p ../public`;
	await $`curl -L ${downloadUrl} -o ${targetAsset}`;
	await $`unzip -p ${targetAsset} index.html > ../public/index.html`;
	await $`rm ${targetAsset}`;
};

const json = await getLatestRelease();
const targetAsset = "AllInOne.zip";
const downloadUrl = getDownloadUrl(targetAsset);
await downloadAndExtract(targetAsset, downloadUrl);
