package kube

objByKind: service: funquoter: {
	apiVersion: "v1"
	metadata: {
		name: "funquoter"
	}
	spec: {
		ports: [{
			port:       80
			targetPort: 3000
		}]
	}
}
objByKind: deployment: funquoter: {
	apiVersion: "apps/v1"
	metadata: name: "funquoter"
	spec: {
		replicas: 1
		selector: matchLabels: app: "funquoter"
		template: {
			spec: containers: [{
				name:  "application"
				image: "k3d-registry.acme.com:5000/fosdem2022/funquoter"
			}]
		}
	}
}
