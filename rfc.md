Title: CI/CD Feedback Standard: Integration of CI/CD into Forges

1. Introduction

   This document proposes the "CI/CD Feedback Standard" for integrating
   Continuous Integration/Continuous Deployment (CI/CD) engines with
   Version Control Systems (VCS). The standard aims to improve
   interoperability and provide a consistent interface for CI/CD
   information across different platforms.

2. Standard Components

   The proposed standard consists of three main components:

   2.1. Well-Known URL
   2.2. Optional HTTP Header
   2.3. JSON Schema for Pipeline Information

3. Well-Known URL

   CI/CD engines implementing this standard MUST provide a well-known URL:

     /.well-known/cicd-feedback

   This URL, when accessed via HTTP GET, MUST return a JSON response
   conforming to the schema defined in the "well-known_schema.json"
   file. The response provides basic information about the CI/CD engine's
   capabilities and support for this standard.

4. Optional HTTP Header

   When a forge sends a webhook to a CI/CD engine, the engine MAY include
   an optional HTTP header in its response to signal support for this
   standard. The header MUST contain:

   - An HTTPS resource URL to retrieve information about the invoked pipeline
   - Optionally, an authorization header named "CICD-Authorization" which
     the forge MUST use when retrieving pipeline information

   Example:
     CICD-Feedback: https://cicd.example.com/pipeline/123
     CICD-Authorization: Bearer <token>

5. JSON Schema for Pipeline Information

   The HTTPS resource URL provided in the HTTP header MUST return a JSON
   response conforming to the schema defined in the "feedback_schema.json"
   file. The schema includes basic pipeline information (ID, status, etc.),
   whether manual action is required, the generated logs and artifacts,
   and the workflows and steps involved in the pipeline.
   The schema also includes an "external URI" field that provides a link
   to the CI/CD engine's native interface. This URI MUST be displayed if
   "requiresManualAction" is true.

   The forge SHOULD refetch the pipeline information periodically based on
   the pipeline status. The forge determines the appropriate interval for
   refetching the data, considering factors such as the current status of
   the pipeline and the likelihood of the status changing over time.

6. Security Considerations

   Implementations MUST use HTTPS for all URLs specified in this standard
   to ensure secure communication. The optional "CICD-Authorization" header
   provides a mechanism for restricting access to sensitive pipeline
   information, ensuring that only authorized parties can retrieve the details.

7. Conclusion

   The "CI/CD Feedback Standard" aims to establish a unified interface
   for CI/CD engines to interact with VCS, enhancing interoperability and
   user experience across different platforms. By adopting this standard,
   CI/CD engines and VCS can work together more effectively, providing
   developers with a streamlined and integrated CI/CD workflow.
