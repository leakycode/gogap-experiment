#import "GoURLProtocol.h"
#import "_cgo_export.h"

@implementation GoURLProtocol

+ (BOOL)canInitWithRequest:(NSURLRequest *)request {
    return [@"go" isEqualToString: request.URL.scheme];
}

+ (NSURLRequest *)canonicalRequestForRequest:(NSURLRequest *)request {
    return request;
}

+ (BOOL)requestIsCacheEquivalent:(NSURLRequest *)a toRequest:(NSURLRequest *)b {
    return [super requestIsCacheEquivalent:a toRequest:b];
}

- (void)startLoading {
    char * urlString = [[self.request.URL description] cStringUsingEncoding:NSUTF8StringEncoding];

    char buffer[10000];
    int bufSize = sizeof(buffer);
    handleURL(urlString, buffer, &bufSize);

    NSData* data = [NSData dataWithBytes:buffer length:bufSize];
    
    NSURLResponse *response = [[NSURLResponse alloc] initWithURL:self.request.URL MIMEType:@"text/html" expectedContentLength:data.length textEncodingName:@"utf8"];
    [[self client] URLProtocol:self didReceiveResponse:response cacheStoragePolicy:NSURLCacheStorageAllowed];

    [[self client] URLProtocol:self didLoadData:data];
    [[self client] URLProtocolDidFinishLoading:self];

    //self.connection = [NSURLConnection connectionWithRequest:self.request delegate:self];
}

- (void)stopLoading {
//    [self.connection cancel];
  //  self.connection = nil;
}

@end
